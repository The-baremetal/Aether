package networking
import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3" // HTTP/3 support
	"golang.org/x/net/http2"           // HTTP/2 support
	"updator/progressbar"                 // Custom progress bar
)
const (
	tempDir = "flux_update_temp"
	barWidth = 40
	desiredUDPBufferSize = 7168 * 1024
)
func DownloadFile(url string, numParts int, outputPath string) error {
	fmt.Printf("Blaze downloader by luohoa97\nThis updator isn't based on GIT, things may break\nbtw run in sudo if it isn't in sudo lol")
	quicTransport := &http3.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		QUICConfig: &quic.Config{
			Allow0RTT: true,
		},
	}
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
if err == nil {
    udpConn.SetReadBuffer(desiredUDPBufferSize)
    udpConn.SetWriteBuffer(desiredUDPBufferSize)
    quicTransport.Dial = func(ctx context.Context, addr string, tlsCfg *tls.Config, cfg *quic.Config) (quic.EarlyConnection, error) {
        remoteAddr, err := net.ResolveUDPAddr("udp", addr)
        if err != nil {
            return nil, fmt.Errorf("failed to resolve address: %w", err)
        }
        return quic.DialEarly(ctx, udpConn, remoteAddr, tlsCfg, cfg)
    }
} else {
    fmt.Printf("Warning: Failed to increase UDP buffer size: %v\n", err)
}
	httpTransport := &http.Transport{}
	http2.ConfigureTransport(httpTransport)
	client := &http.Client{Transport: quicTransport}
	resp, err := client.Head(url)
	if err != nil {
		fmt.Println("Falling back to HTTP/1.1 or HTTP/2:", err)
		client.Transport = httpTransport
		resp, err = client.Head(url)
		if err != nil {
			return fmt.Errorf("error getting file info: %v", err)
		}
	}
	defer resp.Body.Close()
	totalSize := resp.ContentLength
	fmt.Printf("Total file size: %d bytes\n", totalSize)
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()
	pb := progressbar.NewProgressBar(barWidth, "Downloading")
	var wg sync.WaitGroup
	chunkSize := int64(math.Ceil(float64(totalSize) / float64(numParts)))
	var totalBytesDownloaded int64
	var totalTimeElapsed time.Duration
	var pbMutex sync.Mutex
	for i := 0; i < numParts; i++ {
		wg.Add(1)
		go downloadChunk(url, file, i, chunkSize, totalSize, &wg, pb, &totalBytesDownloaded, &totalTimeElapsed, &pbMutex, client)
	}
	wg.Wait()
	return nil
}
func downloadChunk(url string, file *os.File, partIndex int, chunkSize, totalSize int64, wg *sync.WaitGroup, pb *progressbar.ProgressBar, totalBytesDownloaded *int64, totalTimeElapsed *time.Duration, pbMutex *sync.Mutex, client *http.Client) {
	defer wg.Done()
	start := int64(partIndex) * chunkSize
	end := start + chunkSize - 1
	if end >= totalSize {
		end = totalSize - 1
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error downloading chunk:", err)
		return
	}
	defer resp.Body.Close()
	startTime := time.Now()
	_, err = file.Seek(start, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}
	buffer := make([]byte, 1024*1024)
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading chunk:", err)
			return
		}
		if n == 0 {
			break
		}
		_, err = file.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing chunk:", err)
			return
		}
		*totalBytesDownloaded += int64(n)
		*totalTimeElapsed = time.Since(startTime)
		currentSpeed := float64(*totalBytesDownloaded) / (*totalTimeElapsed).Seconds() / 1024 / 1024
		remainingBytes := totalSize - *totalBytesDownloaded
		etaSeconds := float64(remainingBytes) / (currentSpeed * 1024 * 1024)
		eta := time.Duration(etaSeconds * float64(time.Second))
		rawPercentage := float64(*totalBytesDownloaded) / float64(totalSize) * 100
		pbMutex.Lock()
		pb.UpdateProgress(rawPercentage, fmt.Sprintf("Speed: %.2f MB/s ETA: %s", currentSpeed, eta))
		pbMutex.Unlock()
	}
}

func NewFallbackClient() *http.Client {
    quicTransport := &http3.RoundTripper{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    httpTransport := &http.Transport{
        Proxy: http.ProxyFromEnvironment,
        DialContext: (&net.Dialer{
            Timeout:   30 * time.Second,
            KeepAlive: 30 * time.Second,
        }).DialContext,
        MaxIdleConns:          100,
        IdleConnTimeout:       90 * time.Second,
        TLSHandshakeTimeout:   10 * time.Second,
        ExpectContinueTimeout: 1 * time.Second,
    }
    http2.ConfigureTransport(httpTransport)
    
    return &http.Client{
        Timeout: 60 * time.Second,
        Transport: quicTransport,
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            if len(via) >= 10 {
                return fmt.Errorf("stopped after 10 redirects")
            }
            return nil
        },
    }
}

func ReplaceExistingFiles() error {
    return filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if path == tempDir {
            return nil
        }

        relPath, err := filepath.Rel(tempDir, path)
        if err != nil {
            return err
        }

        targetPath := filepath.Join(".", relPath)

        if info.IsDir() {
            return os.MkdirAll(targetPath, 0755)
        }
        input, err := os.ReadFile(path)
        if err != nil {
            return err
        }
        err = os.WriteFile(targetPath, input, info.Mode())
        if err != nil {
            return err
        }

        return nil
    })
}