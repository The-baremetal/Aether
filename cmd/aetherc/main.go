package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"FLUX/parser"
	"FLUX/src/aether/ast"
	"FLUX/src/aether/visitor"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}
	fileName := os.Args[1]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	input := string(content)
	is := antlr.NewInputStream(input)
	lexer := parser.NewLua_grammar_antlr4Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	antlrParser := parser.NewLua_grammar_antlr4Parser(stream)
	visitor := visitor.NewCustomVisitor(antlrParser)
	tree := antlrParser.Program()
	astNode := visitor.Visit(tree)
	fmt.Println("AST:")
	if programNode, ok := astNode.(*ast.ProgramNode); ok {
		fmt.Println(programNode.String(0))
	}
}
