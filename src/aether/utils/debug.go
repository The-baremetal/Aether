package utils

import (
	"fmt"
	"os"
)

var (
	Verbose     bool
	ProgressBar PrintLiner
)

type PrintLiner interface {
	PrintLine(text string)
}

func SetFlagDebug(v bool) {
	Verbose = v
}

func DebugPrintf(format string, a ...interface{}) {
	if Verbose {
		if ProgressBar != nil {
			ProgressBar.PrintLine(fmt.Sprintf(format, a...))
		} else {
			fmt.Fprintf(os.Stderr, format, a...)
		}
	}
}

func DebugPrintln(a ...interface{}) {
	if Verbose {
		if ProgressBar != nil {
			ProgressBar.PrintLine(fmt.Sprint(a...))
		} else {
			fmt.Fprint(os.Stderr, a...)
		}
	}
}
