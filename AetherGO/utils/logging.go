package utils

import (
	"fmt"
	"runtime"
)

const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
)

type Theme struct {
	ErrorColor   string
	SuccessColor string
	WarningColor string
	InfoColor    string
	DebugColor   string
	ErrorSymbol  string
	SuccessSymbol string
	WarningSymbol string
	InfoSymbol    string
	DebugSymbol   string
}

var currentTheme = Theme{
	ErrorColor:    ColorRed,
	SuccessColor:  ColorGreen,
	WarningColor:  ColorYellow,
	InfoColor:     ColorCyan,
	DebugColor:    ColorMagenta,
	ErrorSymbol:   "✗ ",
	SuccessSymbol: "✓ ",
	WarningSymbol: "⚠ ",
	InfoSymbol:    "ℹ ",
	DebugSymbol:   "🐛 ",
}

func SetTheme(newTheme Theme) {
	currentTheme = newTheme
}

func colorPrintf(colorCode string, symbol string, format string, args ...interface{}) {
	fullPrefix := colorCode + symbol
	if runtime.GOOS != "plan9" {
		fmt.Print(fullPrefix)
	} else {
		fmt.Print(symbol)
	}
	fmt.Printf(format, args...)
	if runtime.GOOS != "plan9" {
		fmt.Print(ColorReset)
	}
}

func LogError(format string, args ...interface{}) {
	colorPrintf(currentTheme.ErrorColor, currentTheme.ErrorSymbol, format+"\n", args...)
}

func LogSuccess(format string, args ...interface{}) {
	colorPrintf(currentTheme.SuccessColor, currentTheme.SuccessSymbol, format+"\n", args...)
}

func LogWarning(format string, args ...interface{}) {
	colorPrintf(currentTheme.WarningColor, currentTheme.WarningSymbol, format+"\n", args...)
}

func LogInfo(format string, args ...interface{}) {
	colorPrintf(currentTheme.InfoColor, currentTheme.InfoSymbol, format+"\n", args...)
}

func LogDebug(format string, args ...interface{}) {
	colorPrintf(currentTheme.DebugColor, currentTheme.DebugSymbol, format+"\n", args...)
} 