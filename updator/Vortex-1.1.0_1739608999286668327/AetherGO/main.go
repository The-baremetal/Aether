package main

import (
	"FLUX/AetherGO/core"
	"FLUX/AetherGO/utils"
	"fmt"
	"os"
	"runtime"
)

func main() {
	setTerminalTitle("Nulsis")
	defer resetTerminalColors()
	if err := run(); err != nil {
		utils.LogError("Error: %v", err)
		os.Exit(1)
	}
}

const (
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
)

func setTerminalTitle(title string) {
	if runtime.GOOS != "plan9" {
		fmt.Printf("\033]0;%s\a", title)
		fmt.Print("\033[36m")
	}
}

func colorPrintf(colorCode string, format string, args ...interface{}) {
	if runtime.GOOS != "plan9" {
		fmt.Print(colorCode)
	}
	fmt.Printf(format, args...)
	if runtime.GOOS != "plan9" {
		fmt.Print(colorReset)
	}
}

// Standard color coding functions
func logSuccess(format string, args ...interface{}) {
	colorPrintf(colorGreen+"✓ ", format, args...)
}

func logError(format string, args ...interface{}) {
	colorPrintf(colorRed+"✗ ", format, args...)
}

func logWarning(format string, args ...interface{}) {
	colorPrintf(colorYellow+"⚠ ", format, args...)
}

func logInfo(format string, args ...interface{}) {
	colorPrintf(colorCyan+"ℹ ", format, args...)
}

func logDebug(format string, args ...interface{}) {
	colorPrintf(colorMagenta+"🐛 ", format, args...)
}

func resetTerminalColors() {
	if runtime.GOOS != "plan9" {
		fmt.Print("\033[0m")
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: %s <file>", os.Args[0])
	}
	
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return fmt.Errorf("file read error: %w", err)
	}

	compiler := core.NewCompiler()
	if err := compiler.SetupParser(string(content)); err != nil {
		return err
	}

	if err := compiler.GenerateIR(); err != nil {
		return err
	}

	if err := compiler.ProcessAST(); err != nil {
		return fmt.Errorf("AST processing failed: %w", err)
	}

	irOutput, err := compiler.Finalize()
	if err != nil {
		return fmt.Errorf("finalization failed: %w", err)
	}

	if err := os.WriteFile("output.ll", []byte(irOutput), 0644); err != nil {
		return fmt.Errorf("failed to write IR file: %w", err)
	}

	logSuccess("Compilation successful!")
	return nil
}