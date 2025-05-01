package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"FLUX/src/aether/lexer"
	"FLUX/src/aether/parser"
	"FLUX/src/aether/utils"
	"FLUX/src/lib/progressbar" // ADDED
	"github.com/gookit/color" // ADDED
)

const Copyright = `Copyright (C) 2021 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

Learn more <https://www.gnu.org/licenses/>`

func main() {
	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	Target := buildCmd.String("target", "", "Target architecture (e.g., x86-64, ARM)")
	OptLevel := buildCmd.Int("opt", 0, "Optimization level (0-3)")
	output := buildCmd.String("o", "", "Output file")
	verbose := buildCmd.Bool("v", false, "Enable verbose debugging output")
	version := buildCmd.Bool("version", false, "Show version information")

	_ = Target
	_ = OptLevel
	_ = version

	// Define colors for output
	red := color.FgRed.Render
	yellow := color.FgYellow.Render
	white := color.FgWhite.Render

	if len(os.Args) < 2 {
		fmt.Println("😿 Command syntax error, there are not enough arguments in your command. Try running help to get back on track😸")
		os.Exit(1)
	}

	utils.SetFlagDebug(*verbose)

	switch os.Args[1] {
	case "version":
		content, err := ioutil.ReadFile("currentversion.vortex")
		if err != nil {
			fmt.Println("Error reading version file:", err)
			os.Exit(1)
		}
		fmt.Println(string(content))
		os.Exit(0)
	case "license", "agreement", "clause":
		fmt.Println(string(Copyright))
		os.Exit(0)
	case "build":
		buildCmd.Parse(os.Args[2:])

		if len(buildCmd.Args()) == 0 {
			fmt.Println("😿 Command syntax error, you have not supplied the required build arguments in your command. Try running help to get back on track😸")
			os.Exit(1)
		}

		filename := buildCmd.Args()[0]
		source, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}

		utils.DebugPrintf("Successfully read file: %s\n", filename)
		bar := progressbar.NewProgressBar(100, "Lexing")

		l := lexer.NewLexer(string(source))
		bar.UpdateProgress(25, "Lexing")
		outputPath := *output
		if outputPath == "" {
			ext := filepath.Ext(filename)
			outputPath = filename[:len(filename)-len(ext)] + ".json"
		}

		outputExt := filepath.Ext(outputPath)

		var outputData []byte
		var barMessage string
		var jsonOutput []byte
		var numErrors int = 0 // Define numErrors here

		switch outputExt {
		case ".lex":
			var tokens []lexer.Token
			for {
				tok := l.NextToken()
				if tok.Type == lexer.EOF {
					break
				}
				tokens = append(tokens, tok)
			}
			jsonOutput, err = json.MarshalIndent(tokens, "", "  ")
			if err != nil {
				fmt.Printf("Error marshalling JSON: %s\n", err)
				os.Exit(1)
			}
			outputData = jsonOutput
			barMessage = "Lexing"
			bar.UpdateProgress(75, "Marshalling")
			err = ioutil.WriteFile(outputPath, outputData, 0644)
				if err != nil {
					fmt.Printf("Error writing to file: %s\n", err)
					os.Exit(1)
				}
		case ".par":
			p := parser.NewParser(l)
			program := p.ParseProgram()
			errors := p.Errors()
			numErrors = len(errors)

			if numErrors > 0 {
				fmt.Println(red("😿 Parsing failed with the following errors:"))
				for _, err := range errors {
					fmt.Println(red(err))
				}
			}

			if numErrors == 0 {
				jsonOutput, err = json.MarshalIndent(program, "", "  ")
				if err != nil {
					fmt.Printf("Error marshalling JSON: %s\n", err)
					os.Exit(1)
				}
				outputData = jsonOutput
				barMessage = "Parsing"
				bar.UpdateProgress(75, "Marshalling")

				err = ioutil.WriteFile(outputPath, outputData, 0644)
				if err != nil {
					fmt.Printf("Error writing to file: %s\n", err)
					os.Exit(1)
				}

				utils.DebugPrintf("JSON tokens generated to: %s\n", outputPath)
			} else {
				barMessage = "Parsing"
				outputData = []byte{}
			}
		default:
			var tokens []lexer.Token
			for {
				tok := l.NextToken()
				if tok.Type == lexer.EOF {
					break
				}
				tokens = append(tokens, tok)
			}
			jsonOutput, err = json.MarshalIndent(tokens, "", "  ")
			if err != nil {
				fmt.Printf("Error marshalling JSON: %s\n", err)
				os.Exit(1)
			}
			outputData = jsonOutput
			barMessage = "Lexing"
			bar.UpdateProgress(75, "Marshalling")
		}

		bar.UpdateProgress(100)
		bar.Finish()

		// Print compilation status
		errorColor := white
		if numErrors > 0 {
			errorColor = red
		}
		warningColor := white
		numWarnings := 0 // TODO: Implement warnings
		if numWarnings > 0 {
			warningColor = yellow
		}

		var emoji string
		status := white("successful")
		emoji = "😺" // Grinning cat for success
		if numErrors > 0 {
			status = red("failed")
			emoji = "😿" // Crying cat for errors
		} else if numWarnings > 0 {
			emoji = "😢" // Cry emoji for warnings
		}

		fmt.Printf("%s %s %s with %s errors and %s warnings\n",
			emoji,
			barMessage,
			status,
			errorColor(fmt.Sprintf("%d", numErrors)),
			warningColor(fmt.Sprintf("%d", numWarnings)),
		)
	case "run":
		fmt.Println("Run command not implemented yet")
		os.Exit(1)
	case "help":
		fmt.Println("Usage: aetherc <command> [options] 😺")
		fmt.Println("Commands:")
		fmt.Println("  build: Build a file 🔨")
		fmt.Println("  run: Run a file (not implemented yet) 👟")
		os.Exit(1)
	default:
		fmt.Println("😿 Command syntax error, invalid command. Try running help to get back on track😸", os.Args[1])
		os.Exit(1)
	}
}
