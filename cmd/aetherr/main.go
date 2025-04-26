package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"FLUX/src/aether/utils"
)

func main() {
	interpType := flag.String("interp", "jit", "Interpreter type (jit or aot)")
	verbose := flag.Bool("v", false, "Enable verbose debugging output")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: aetherr [options] <file>")
		os.Exit(1)
	}

	utils.Verbose = *verbose

	filename := flag.Args()[0]
	utils.DebugPrintf("Running file: %s with interpreter: %s\n", filename, *interpType)

	code, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	utils.DebugPrintf("Successfully read file: %s\n", filename)

	var cmd *exec.Cmd
	switch *interpType {
	case "jit":
		cmd = exec.Command("lli")
	case "aot":
		fmt.Println("AOT compilation not implemented yet")
		os.Exit(1)
	default:
		fmt.Println("Invalid interpreter type:", *interpType)
		os.Exit(1)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Pass the LLVM code to lli's stdin
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
		os.Exit(1)
	}

	go func() {
		defer stdin.Close()
		stdin.Write(code)
		utils.DebugPrintln("Successfully wrote code to lli stdin")
	}()

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing LLVM code:", err)
		os.Exit(1)
	}

	utils.DebugPrintln("LLVM code executed successfully")
}
