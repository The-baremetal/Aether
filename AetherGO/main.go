package main

import (
	"FLUX/AetherGO/core"
	"FLUX/AetherGO/utils"
	"os"
	"fmt"
	"runtime"
)

func main() {
	setTerminalTitle("Nulsis (Aether compiler codename)")
	defer resetTerminalColors()
	
	if len(os.Args) < 2 {
		panic("Please provide source file")
	}

	sourceCode, _ := os.ReadFile(os.Args[1])
	ast := core.ParseSource(string(sourceCode))
	compiler := core.NewCompiler()
	if err := compiler.Compile(ast); err != nil {
		panic(err)
	}
	println(compiler.Module.String())
	println("smth", ast)
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

func readSource(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		utils.LogError("File read error: %v", err)
		os.Exit(1)
	}
	return string(content)
}