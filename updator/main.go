package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
	"net/http"
	"archive/tar"
	"compress/gzip"
	"strings"
	"strconv"
	"encoding/json"
	"time"

	"updator/networking"
)

const (
	versionFile    = "currentversion.vortex"
	repoURL        = "https://raw.githubusercontent.com/The-baremetal/FLUXASSEMBLY/main/"
	versionPrefix  = "Vortex-"
	repoOwner      = "The-baremetal"
	repoName       = "FLUXASSEMBLY"
	githubAPI      = "https://api.github.com/repos/%s/%s/releases/latest"
	tarballURL     = "https://github.com/%s/%s/archive/refs/tags/%s.tar.gz"
	tempDir        = "flux_update_temp"
)

type VortexVersion struct {
	Core       int
	Flow       int
	Patch      int
	Timestamp  string
	CustomTags string
}

type ReleaseInfo struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func main() {
	localVersion, err := getLocalVersion()
	if err != nil {
		fmt.Printf("Error getting local version: %v\n", err)
		os.Exit(1)
	}

	releaseInfo, err := getLatestReleaseInfo()
	if err != nil {
		fmt.Printf("Error fetching release info: %v\n", err)
		os.Exit(1)
	}

	remoteVersion, err := parseVersionFromTag(releaseInfo.TagName)
	if err != nil {
		fmt.Printf("Error parsing release tag: %v\n", err)
		os.Exit(1)
	}

	if isNewerVersion(remoteVersion, localVersion) {
		fmt.Printf("New version %s available, updating...\n", releaseInfo.TagName)
		
		if err := downloadAndExtractRelease(releaseInfo.TagName); err != nil {
			fmt.Printf("Update failed: %v\n", err)
			cleanupTempFiles()
			os.Exit(1)
		}
		
		fmt.Printf("Successfully updated to version: %s\n", releaseInfo.TagName)
		cleanupTempFiles()
	} else {
		fmt.Println("Already running latest version")
	}
}

func getLocalVersion() (*VortexVersion, error) {
	if _, err := os.Stat(versionFile); os.IsNotExist(err) {
		return &VortexVersion{0, 0, 0, "", ""}, nil
	}
	return parseVersionFromFile(versionFile)
}

func getLatestReleaseInfo() (*ReleaseInfo, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(githubAPI, repoOwner, repoName), nil)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Add("Authorization", "token "+token)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusForbidden && resp.Header.Get("X-RateLimit-Remaining") == "0" {
		resetTime := resp.Header.Get("X-RateLimit-Reset")
		return nil, fmt.Errorf("GitHub API rate limit exceeded. Resets at %s", resetTime)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var release ReleaseInfo
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}
	return &release, nil
}

func downloadAndExtractRelease(tag string) error {
	if err := os.Mkdir(tempDir, 0755); err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	tarballPath := filepath.Join(tempDir, "release.tar.gz")
	url := fmt.Sprintf(tarballURL, repoOwner, repoName, tag)
	if err := networking.DownloadFile(url, 4, tarballPath); err != nil {
		return fmt.Errorf("failed to download release: %w", err)
	}
	if err := extractTarGz(tarballPath); err != nil {
		return fmt.Errorf("failed to extract release: %w", err)
	}
	return networking.ReplaceExistingFiles()
}

func extractTarGz(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		targetPath := filepath.Join(tempDir, strings.TrimPrefix(header.Name, repoName+"-"))
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}

		if header.Typeflag == tar.TypeReg {
			outFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}
	return nil
}

func parseVersionFromTag(tag string) (*VortexVersion, error) {
	cleanTag := strings.TrimPrefix(tag, "v")
	parts := strings.Split(cleanTag, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid tag format")
	}

	core, _ := strconv.Atoi(parts[0])
	flow, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	return &VortexVersion{
		Core:  core,
		Flow:  flow,
		Patch: patch,
	}, nil
}

func parseVersionFromFile(filename string) (*VortexVersion, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading version file: %w", err)
	}

	return parseVersionString(strings.TrimSpace(string(data)))
}

func parseVersionString(versionStr string) (*VortexVersion, error) {
	if !strings.HasPrefix(versionStr, versionPrefix) {
		return nil, fmt.Errorf("invalid version format")
	}

	cleanStr := strings.TrimPrefix(versionStr, versionPrefix)
	parts := strings.SplitN(cleanStr, "_", 2)
	versionPart := parts[0]
	var timestamp, customTags string

	if len(parts) > 1 {
		metadata := strings.SplitN(parts[1], "_", 2)
		timestamp = metadata[0]
		if len(metadata) > 1 {
			customTags = "_" + metadata[1]
		}
	}

	numbers := strings.Split(versionPart, ".")
	if len(numbers) != 3 {
		return nil, fmt.Errorf("invalid version components")
	}

	core, err := strconv.Atoi(numbers[0])
	if err != nil {
		return nil, fmt.Errorf("invalid core version")
	}

	flow, err := strconv.Atoi(numbers[1])
	if err != nil {
		return nil, fmt.Errorf("invalid flow version")
	}

	patch, err := strconv.Atoi(numbers[2])
	if err != nil {
		return nil, fmt.Errorf("invalid patch version")
	}

	return &VortexVersion{
		Core:       core,
		Flow:       flow,
		Patch:      patch,
		Timestamp:  timestamp,
		CustomTags: customTags,
	}, nil
}

func isNewerVersion(remote, local *VortexVersion) bool {
	if remote.Core > local.Core {
		return true
	}
	if remote.Core == local.Core {
		if remote.Flow > local.Flow {
			return true
		}
		if remote.Flow == local.Flow && remote.Patch > local.Patch {
			return true
		}
	}
	return false
}

func formatVersion(v *VortexVersion) string {
	return fmt.Sprintf("%s%d.%d.%d_%s%s", 
		versionPrefix,
		v.Core,
		v.Flow,
		v.Patch,
		v.Timestamp,
		v.CustomTags,
	)
}

func cleanupTempFiles() {
	os.RemoveAll(tempDir)
}
