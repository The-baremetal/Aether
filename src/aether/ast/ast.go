package ast

import (
	"fmt"
	"strings"
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

func (a AssignStatementNode) statementNode() {}

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

func (e ExpressionStatementNode) statementNode() {}

func (e *ExpressionStatementNode) String(indent int) string {
	return fmt.Sprintf("%sExprStmt: %s\n", indentStr(indent), e.Expr.String(0))
}

// ReturnStatementNode represents a return statement.
type ReturnStatementNode struct {
	Values []Expression
}

func (r ReturnStatementNode) statementNode() {}

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

func (i IfStatementNode) statementNode() {}

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
    Value interface{} // Can hold int, float64, string, or table
    Type  string      // Type can be "int", "float", "string", or table
}

func (l *LiteralNode) expressionNode() {}

func (l *LiteralNode) String(indent int) string {
    return fmt.Sprintf("%sLiteral: %v (%s)", indentStr(indent), l.Value, l.Type)
}

// VariableNode represents a variable with a dynamic type.
type VariableNode struct {
	Name    string
	Type    string // This can be "int", "float", "string", "table"
	Mutable bool
}

func (v *VariableNode) expressionNode() {}

func (v *VariableNode) String(indent int) string {
    return fmt.Sprintf("%sVariable: %s (%s)", indentStr(indent), v.Name, v.Type)
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

func indentStr(indent int) string {
	return strings.Repeat("  ", indent)
}

type TypeSpec struct {
	Name string
}

func (ts TypeSpec) expressionNode() {}

type MatchArm struct {
	Pattern Expression
	When    Expression
	Result  Expression
}

func (ma MatchArm) expressionNode() {}

type BlockStmt struct {
	Statements []Statement
}

func (bs BlockStmt) expressionNode() {}
