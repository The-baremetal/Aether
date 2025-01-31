package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"io/ioutil"
	"FLUX/parser"
)

// AST Node types
type NodeType int

const (
	ProgramNode NodeType = iota
	FunctionDecl
	VariableDecl
	BinaryExpr
	CallExpr
	NumberLiteral
	StringLiteral
	Identifier
	// Add more node types as needed
)

type ASTNode struct {
	Type     NodeType
	Value    string
	Children []*ASTNode
}

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
	lexer := parser.Newlua_grammar_antlr4Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.Newlua_grammar_antlr4Parser(stream)
	tree := parser.Program()

	// Create AST builder
	astBuilder := NewASTBuilder()
	
	// Walk the tree with AST listener
	antlr.ParseTreeWalkerDefault.Walk(astBuilder, tree)

	// Get root AST node
	astRoot := astBuilder.GetAST()
	
	// Print AST
	fmt.Println("Abstract Syntax Tree:")
	printAST(astRoot, 0)
}

type ASTBuilder struct {
	*parser.Baselua_grammar_antlr4Listener
	stack []*ASTNode
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{
		stack: make([]*ASTNode, 0),
	}
}

func (b *ASTBuilder) GetAST() *ASTNode {
	return b.stack[0]
}

func (b *ASTBuilder) pushNode(node *ASTNode) {
	b.stack = append(b.stack, node)
}

func (b *ASTBuilder) popNode() *ASTNode {
	if len(b.stack) < 1 {
		panic("stack is empty")
	}
	node := b.stack[len(b.stack)-1]
	b.stack = b.stack[:len(b.stack)-1]
	return node
}

// Example rule transformation
func (b *ASTBuilder) ExitExpression(ctx *parser.ExpressionContext) {
	if ctx.GetChildCount() == 3 { // Binary operation
		right := b.popNode()
		op := ctx.GetChild(1).GetPayload().(*antlr.CommonToken).GetText()
		left := b.popNode()
		
		b.pushNode(&ASTNode{
			Type:  BinaryExpr,
			Value: op,
			Children: []*ASTNode{left, right},
		})
	}
}

func (b *ASTBuilder) ExitNumberLiteral(ctx *parser.NumberLiteralContext) {
	b.pushNode(&ASTNode{
		Type:  NumberLiteral,
		Value: ctx.GetText(),
	})
}

// Add more Exit methods for other rules...

func printAST(node *ASTNode, indent int) {
	fmt.Printf("%s%s", getIndent(indent), node.Type)
	if node.Value != "" {
		fmt.Printf(" (%s)", node.Value)
	}
	fmt.Println()
	
	for _, child := range node.Children {
		printAST(child, indent+2)
	}
}

func getIndent(indent int) string {
	return fmt.Sprintf("%*s", indent, "")
}
