package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"FLUX/parser"
)

//
//===----------------------------------------------------------------------===//
// AST Node Definitions
//===----------------------------------------------------------------------===//

// Node is the base interface for all AST nodes.
type Node interface {
	String(indent int) string
}

// Statement is the interface for statement nodes.
type Statement interface {
	Node
	statementNode() // marker method
}

// Expression is the interface for expression nodes.
type Expression interface {
	Node
	expressionNode() // marker method
}

// ProgramNode represents the entire program.
type ProgramNode struct {
	Statements []Statement
}

func (p *ProgramNode) String(indent int) string {
	var sb strings.Builder
	sb.WriteString(indentStr(indent) + "Program:\n")
	for _, stmt := range p.Statements {
		sb.WriteString(stmt.String(indent + 1))
	}
	return sb.String()
}

// ---------- Statement Nodes ----------

// AssignStatementNode represents an assignment statement.
type AssignStatementNode struct {
	Variable   Expression // left-hand side (variable or lvalue)
	Operator   string
	Expression Expression // right-hand side
}

func (a *AssignStatementNode) statementNode() {}

func (a *AssignStatementNode) String(indent int) string {
	return fmt.Sprintf("%sAssign: (%s %s %s)\n",
		indentStr(indent),
		a.Variable.String(0),
		a.Operator,
		a.Expression.String(0))
}

// ExpressionStatementNode wraps an expression as a statement.
type ExpressionStatementNode struct {
	Expr Expression
}

func (e *ExpressionStatementNode) statementNode() {}

func (e *ExpressionStatementNode) String(indent int) string {
	return fmt.Sprintf("%sExprStmt: %s\n", indentStr(indent), e.Expr.String(0))
}

// ReturnStatementNode represents a return statement.
type ReturnStatementNode struct {
	Values []Expression
}

func (r *ReturnStatementNode) statementNode() {}

func (r *ReturnStatementNode) String(indent int) string {
	var sb strings.Builder
	sb.WriteString(indentStr(indent) + "Return:\n")
	for _, v := range r.Values {
		sb.WriteString(v.String(indent + 1))
	}
	return sb.String()
}

// IfStatementNode represents an if statement.
type IfStatementNode struct {
	Condition Expression
	Then      *BlockNode
	Else      *BlockNode // may be nil
}

func (i *IfStatementNode) statementNode() {}

func (i *IfStatementNode) String(indent int) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%sIf:\n", indentStr(indent)))
	sb.WriteString(fmt.Sprintf("%sCondition: %s\n", indentStr(indent+1), i.Condition.String(0)))
	sb.WriteString(fmt.Sprintf("%sThen:\n%s", indentStr(indent+1), i.Then.String(indent+2)))
	if i.Else != nil {
		sb.WriteString(fmt.Sprintf("%sElse:\n%s", indentStr(indent+1), i.Else.String(indent+2)))
	}
	return sb.String()
}

// BlockNode represents a sequence of statements.
type BlockNode struct {
	Statements []Statement
}

func (b *BlockNode) String(indent int) string {
	var sb strings.Builder
	for _, stmt := range b.Statements {
		sb.WriteString(stmt.String(indent))
	}
	return sb.String()
}

// ---------- Expression Nodes ----------

// BinaryExpressionNode represents a binary expression.
type BinaryExpressionNode struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryExpressionNode) expressionNode() {}

func (b *BinaryExpressionNode) String(indent int) string {
	return fmt.Sprintf("%sBinaryExpr: (%s %s %s)",
		indentStr(indent),
		b.Left.String(0),
		b.Operator,
		b.Right.String(0))
}

// UnaryExpressionNode represents a unary expression.
type UnaryExpressionNode struct {
	Operator string
	Expr     Expression
}

func (u *UnaryExpressionNode) expressionNode() {}

func (u *UnaryExpressionNode) String(indent int) string {
	return fmt.Sprintf("%sUnaryExpr: (%s %s)",
		indentStr(indent),
		u.Operator,
		u.Expr.String(0))
}

// LiteralNode represents a literal.
type LiteralNode struct {
	Value string
}

func (l *LiteralNode) expressionNode() {}

func (l *LiteralNode) String(indent int) string {
	return fmt.Sprintf("%sLiteral: %s", indentStr(indent), l.Value)
}

// VariableNode represents a variable.
type VariableNode struct {
	Name string
}

func (v *VariableNode) expressionNode() {}

func (v *VariableNode) String(indent int) string {
	return fmt.Sprintf("%sVariable: %s", indentStr(indent), v.Name)
}

// FunctionCallNode represents a function call.
type FunctionCallNode struct {
	Callee    Expression
	Arguments []Expression
}

func (f *FunctionCallNode) expressionNode() {}

func (f *FunctionCallNode) String(indent int) string {
	var args []string
	for _, arg := range f.Arguments {
		args = append(args, arg.String(0))
	}
	return fmt.Sprintf("%sFunctionCall: %s(%s)",
		indentStr(indent),
		f.Callee.String(0),
		strings.Join(args, ", "))
}

// generateLLVMCodeForBinaryExpression generates LLVM code for a BinaryExpressionNode.
func generateLLVMCodeForBinaryExpression(b *BinaryExpressionNode) string {
	left := generateLLVMCodeForExpression(b.Left)
	right := generateLLVMCodeForExpression(b.Right)
	switch b.Operator {
	case "+":
		return fmt.Sprintf("add i32 %s, %s", left, right)
	case "-":
		return fmt.Sprintf("sub i32 %s, %s", left, right)
	case "*":
		return fmt.Sprintf("mul i32 %s, %s", left, right)
	case "/":
		return fmt.Sprintf("sdiv i32 %s, %s", left, right)
	default:
		return ""
	}
}
func generateLLVMCodeForExpression(e Expression) string {
	switch v := e.(type) {
	case *LiteralNode:
		return v.Value
	case *VariableNode:
		return "%" + v.Name
	case *BinaryExpressionNode:
		return generateLLVMCodeForBinaryExpression(v)
	case *UnaryExpressionNode:
		return generateLLVMCodeForUnaryOperation(v)
	default:
		return ""
	}
}
func generateLLVMCodeForUnaryOperation(u *UnaryExpressionNode) string {
	expr := generateLLVMCodeForExpression(u.Expr)
	if u.Operator == "-" {
		return fmt.Sprintf("sub i32 0, %s", expr)
	}
	return ""
}
func indentStr(indent int) string {
	return strings.Repeat("  ", indent)
}

//
//===----------------------------------------------------------------------===//
// Visitor Implementation
//===----------------------------------------------------------------------===//
type CustomVisitor struct {
	*parser.BaseLua_grammar_antlr4Visitor
	Parser *parser.Lua_grammar_antlr4Parser
}

func NewCustomVisitor(p *parser.Lua_grammar_antlr4Parser) *CustomVisitor {
	return &CustomVisitor{
		BaseLua_grammar_antlr4Visitor: &parser.BaseLua_grammar_antlr4Visitor{},
		Parser:                        p,
	}
}
func (v *CustomVisitor) visitExpr(tree antlr.ParseTree) Expression {
	res := v.Visit(tree)
	if res == nil {
		return &LiteralNode{Value: tree.GetText()}
	}
	if expr, ok := res.(Expression); ok {
		return expr
	}
	return &LiteralNode{Value: fmt.Sprintf("%v", res)}
}
func (v *CustomVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	stmts := []Statement{}
	// Assuming AllStatement() returns a slice of IStatementContext.
	for _, stmtCtx := range ctx.AllStatement() {
		if pt, ok := stmtCtx.(antlr.ParseTree); ok {
			res := v.Visit(pt)
			if res == nil {
				continue
			}
			if s, ok := res.(Statement); ok {
				stmts = append(stmts, s)
			} else if expr, ok := res.(Expression); ok {
				stmts = append(stmts, &ExpressionStatementNode{Expr: expr})
			}
		}
	}
	return &ProgramNode{Statements: stmts}
}

// VisitAssignStmt corresponds to the assignStatement rule.
// Grammar: assignStatement : variable (assignOp | incrOp | shiftAssignOp | coalesceOp) expression ;
func (v *CustomVisitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	child0 := ctx.GetChild(0).(antlr.ParseTree)
	varNode := v.visitExpr(child0)

	child1 := ctx.GetChild(1)
	op := ""
	if term, ok := child1.(antlr.TerminalNode); ok {
		op = term.GetText()
	}

	child2 := ctx.GetChild(2).(antlr.ParseTree)
	exprNode := v.visitExpr(child2)

	return &AssignStatementNode{
		Variable:   varNode,
		Operator:   op,
		Expression: exprNode,
	}
}

func (v *CustomVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	child := ctx.GetChild(0).(antlr.ParseTree)
	expr := v.visitExpr(child)

	llvmCode := ""
	var finalExpr Expression

	if lit, ok := expr.(*LiteralNode); ok {
		line := lit.Value
		if strings.HasPrefix(line, "local") {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				variableName := strings.TrimSpace(strings.ReplaceAll(parts[0], "local", ""))
				initialValue := strings.TrimSpace(parts[1])
				llvmCode = fmt.Sprintf("%%%s = alloca i32\nstore i32 %s, i32* %%%s\n", variableName, initialValue, variableName)
				fmt.Println(llvmCode)
				finalExpr = &VariableNode{Name: variableName}
			}
		} else {
			var operator string
			if strings.Contains(lit.Value, "+") {
				operator = "+"
			} else if strings.Contains(lit.Value, "-") {
				operator = "-"
			} else if strings.Contains(lit.Value, "*") {
				operator = "*"
			} else if strings.Contains(lit.Value, "/") {
				operator = "/"
			} else if strings.Contains(lit.Value, "^") {
				operator = "^"
			}

			if operator != "" {
				parts := strings.Split(lit.Value, operator)
				if len(parts) == 2 {
					left := &VariableNode{Name: strings.TrimSpace(parts[0])}
					right := &VariableNode{Name: strings.TrimSpace(parts[1])}
					finalExpr = &BinaryExpressionNode{
						Left:     left,
						Operator: operator,
						Right:    right,
					}
				}
			} else {
				finalExpr = expr
			}
		}
	} else {
		finalExpr = expr
	}
	if varNode, ok := finalExpr.(*VariableNode); ok {
		llvmCode = fmt.Sprintf("load i32, i32* %%%s", varNode.Name)
		fmt.Println(llvmCode)
		return nil
	}

	if binExpr, ok := finalExpr.(*BinaryExpressionNode); ok {
		llvmCode = generateLLVMCodeForBinaryExpression(binExpr)
		fmt.Println(llvmCode)
		return nil
	}

	return &ExpressionStatementNode{Expr: expr}
}

// VisitReturnStmt corresponds to the returnStatement rule.
// Grammar: returnStatement : RETURN (expression (',' expression)* | functionCall)? ;
func (v *CustomVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	exprs := []Expression{}
	for i := 1; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if term, ok := child.(antlr.TerminalNode); ok && term.GetText() == "," {
			continue
		}
		if pt, ok := child.(antlr.ParseTree); ok {
			expr := v.visitExpr(pt)
			exprs = append(exprs, expr)
		}
	}
	return &ReturnStatementNode{Values: exprs}
}

// VisitIfStatement corresponds to the if-statement rule.
// Grammar: ifStatement : 'if' expression 'then' block ('elseif' expression 'then' block)* ('else' block)? 'end' ;
func (v *CustomVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	cond := v.visitExpr(ctx.Expression(0).(antlr.ParseTree))
	blocks := ctx.AllBlock()
	thenBlock := &BlockNode{Statements: v.visitBlock(blocks[0])}
	var elseBlock *BlockNode
	if len(blocks) > 1 {
		elseBlock = &BlockNode{Statements: v.visitBlock(blocks[len(blocks)-1])}
	}
	return &IfStatementNode{
		Condition: cond,
		Then:      thenBlock,
		Else:      elseBlock,
	}
}

// Helper: visitBlock collects all statements from a block.
func (v *CustomVisitor) visitBlock(ctx parser.IBlockContext) []Statement {
	stmts := []Statement{}
	for _, sCtx := range ctx.AllStatement() {
		if pt, ok := sCtx.(antlr.ParseTree); ok {
			res := v.Visit(pt)
			if res == nil {
				continue
			}
			if s, ok := res.(Statement); ok {
				stmts = append(stmts, s)
			} else if expr, ok := res.(Expression); ok {
				stmts = append(stmts, &ExpressionStatementNode{Expr: expr})
			}
		}
	}
	return stmts
}

// VisitLiteral corresponds to the literal rule.
func (v *CustomVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return &LiteralNode{Value: ctx.GetText()}
}

// VisitVariable corresponds to the variable rule.
func (v *CustomVisitor) VisitVariable(ctx *parser.VariableContext) interface{} {
	return &VariableNode{Name: ctx.GetText()}
}

// VisitBinaryOperation corresponds to binaryOperation.
func (v *CustomVisitor) VisitBinaryOperation(ctx *parser.BinaryOperationContext) interface{} {
	left := v.visitExpr(ctx.Expression(0).(antlr.ParseTree))
	op := ctx.ArithOp().GetText()
	right := v.visitExpr(ctx.Expression(1).(antlr.ParseTree))
	return &BinaryExpressionNode{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}

// VisitUnaryOperation corresponds to unaryOperation.
func (v *CustomVisitor) VisitUnaryOperation(ctx *parser.UnaryOperationContext) interface{} {
	op := ctx.UnaryOp().GetText()
	expr := v.visitExpr(ctx.Expression().(antlr.ParseTree))
	return &UnaryExpressionNode{
		Operator: op,
		Expr:     expr,
	}
}

// Fallback Visit method.
func (v *CustomVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

//
//===----------------------------------------------------------------------===//
// Main entrypoint
//===----------------------------------------------------------------------===//

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
	visitor := NewCustomVisitor(antlrParser)
	tree := antlrParser.Program()
	ast := visitor.Visit(tree)
	fmt.Println("LLVM Code:")
	if programNode, ok := ast.(*ProgramNode); ok {
		for _, stmt := range programNode.Statements {
			if exprStmt, ok := stmt.(*ExpressionStatementNode); ok {
				fmt.Println(generateLLVMCodeForExpression(exprStmt.Expr))
			}
		}
	}
}
