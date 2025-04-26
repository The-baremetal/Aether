package newparser_test

import (
	"FLUX/src/aether/ast"
	"FLUX/src/aether/parser"
	"FLUX/src/lib/lexer"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		t.Fatalf("Error reading directory: %s", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".lua") && file.Name() != "input.lua" {
			t.Run(file.Name(), func(t *testing.T) {
				input, err := ioutil.ReadFile(file.Name())
				if err != nil {
					t.Fatalf("Error reading file: %s", err)
				}

				l := lexer.NewLexer(string(input))
				p := parser.NewParser(l)
				program := p.ParseProgram()

				if len(p.Errors()) > 0 {
					for _, err := range p.Errors() {
						t.Errorf("Parser error: %s", err)
					}
					t.FailNow()
				}

				if program == nil {
					t.Fatalf("ParseProgram() returned nil")
				}

				fmt.Printf("Parsed %s successfully\n", file.Name())
			})
		}
	}
}

func TestBinaryExpression(t *testing.T) {
	input := `
local a = 1 + 2 * 3
`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser error: %s", err)
		}
		t.FailNow()
	}

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}

	localDeclaration, ok := program.Statements[0].(*ast.LocalDeclaration)
	if !ok {
		t.Fatalf("Expected LocalDeclaration, got %T", program.Statements[0])
	}

	if localDeclaration.Identifier.Name != "a" {
		t.Fatalf("Expected identifier 'a', got %s", localDeclaration.Identifier.Name)
	}
}

func TestConstDeclaration(t *testing.T) {
	input := `
const a = 1 + 2 * 3
`
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		for _, err := range p.Errors() {
			t.Errorf("Parser error: %s", err)
		}
		t.FailNow()
	}

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}

	constDeclaration, ok := program.Statements[0].(*ast.ConstDeclaration)
	if !ok {
		t.Fatalf("Expected ConstDeclaration, got %T", program.Statements[0])
	}

	if constDeclaration.Identifier.Name != "a" {
		t.Fatalf("Expected identifier 'a', got %s", constDeclaration.Identifier.Name)
	}
}
