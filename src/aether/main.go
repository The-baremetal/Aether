package main

import (
	"fmt"
	"FLUX/src/aether/ast"
	"FLUX/src/aether/parser"
	"FLUX/src/lib/lexer"
)

func main() {
	input := `
	local x = 123
	mutable y = 544
	return x
	`

	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	program := p.Parse()

	if program != nil {
		fmt.Println("AST:")
		for _, stmt := range program.Statements {
			fmt.Printf("%T\n", stmt)
		}
		if l.HasErrors() {
			l.PrintErrors()
		}
		if p.errorHandler.HasErrors() {
			p.errorHandler.PrintErrors()
		}
	} else {
		fmt.Println("Parsing failed.")
		if l.HasErrors() {
			l.PrintErrors()
		}
		if p.errorHandler.HasErrors() {
			p.errorHandler.PrintErrors()
		}
	}
}