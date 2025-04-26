package leash_test

import (
	"FLUX/src/aether/ast"
	"FLUX/src/aether/checker"
	"FLUX/src/aether/parser"
	"FLUX/src/lib/lexer"
	"testing"
)

func TestLeashSystem(t *testing.T) {
	input := `
local a = 1
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

	borrowChecker := checker.NewBorrowChecker()
	borrowChecker.Check(localDeclaration)

	if len(borrowChecker.Check(localDeclaration)) != 0 {
		t.Fatalf("Expected 0 errors, got %d", len(borrowChecker.Check(localDeclaration)))
	}

	if localDeclaration.Leash == nil {
		t.Fatalf("Expected leash, got nil")
	}

	if localDeclaration.Leash.GetOwner() != localDeclaration.Expression {
		t.Fatalf("Expected owner to be expression, got %v", localDeclaration.Leash.GetOwner())
	}
}
