package core

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"FLUX/parser"
	"FLUX/AetherGO/types"
)

func ParseSource(code string) *types.ASTNode {
	input := antlr.NewInputStream(code)
	lexer := parser.NewLua_grammar_antlr4Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLua_grammar_antlr4Parser(stream)
	tree := p.Program()
	fmt.Println("RAW PARSE TREE:")
	fmt.Println(tree.ToStringTree([]string{}, p))
	
	visitor := NewCustomVisitor(p)
	result := p.Program().Accept(visitor)
	if result == nil {
		return &types.ASTNode{Type: types.NodeStmt}
	}
	return result.(*types.ASTNode)
}