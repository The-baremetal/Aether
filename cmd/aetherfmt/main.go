package main

/*

import (
	gofmt "fmt"
	"io/ioutil"
	"os"

	"src/aether/ast"
	"src/aether/syntax"
	"src/lib/lexer"
	"src/aether/fmt"
)

func main() {
	// Read input file
	if len(os.Args) < 2 {
		gofmt.Println("Usage: aetherfmt <file>")
		return
	}
	filename := os.Args[1]
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		gofmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Lex and Parse the code
	l := lexer.NewLexer(string(input))
	p := syntax.NewParser(l) // CHANGED
	program := p.ParseProgram()

	// Format the code
	formattedCode := format(program)

	// Print the formatted code
	gofmt.Println(formattedCode)
}

func format(program *ast.Program) string {
	f := fmt.NewFormatter()
	return f.VisitProgram(program)
}
*/
func main(){}
