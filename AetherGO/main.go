package main

import (
	"FLUX/AetherGO/core"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
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

	fmt.Println("Compilation successful!")
	return nil
}