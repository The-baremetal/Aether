package visitor

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"FLUX/parser"
	"FLUX/src/aether/ast"
	"os"
)

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
func (v *CustomVisitor) visitExpr(tree antlr.ParseTree) ast.Expression {
	res := v.Visit(tree)
	if res == nil {
		return &ast.LiteralNode{Value: tree.GetText()}
	}
	if expr, ok := res.(ast.Expression); ok {
		return expr
	}
	return &ast.LiteralNode{Value: fmt.Sprintf("%v", res)}
}
func (v *CustomVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	stmts := []ast.Statement{}
	// Assuming AllStatement() returns a slice of IStatementContext.
	for _, stmtCtx := range ctx.AllStatement() {
		if pt, ok := stmtCtx.(antlr.ParseTree); ok {
			res := v.Visit(pt)
			if res == nil {
				continue
			}
			if s, ok := res.(ast.Statement); ok {
				stmts = append(stmts, s)
			} else if expr, ok := res.(ast.Expression); ok {
				stmts = append(stmts, &ast.ExpressionStatementNode{Expr: expr})
			}
		}
	}
	return &ast.ProgramNode{Statements: stmts}
}

// VisitAssignStmt corresponds to the assignStatement rule.
// Grammar: assignStatement : variable (assignOp | incrOp | shiftAssignOp | coalesceOp) expression ;
func (v *CustomVisitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	if ctx.GetChildCount() < 3 {
		line := ctx.GetStart().GetLine()
		fmt.Printf("Error on line %d: Expected 3 arguments, You did not include all your required arguments in your code.\n", line)
		fmt.Println("Hint: Ensure the format is '[variable] [operator] [expression]'.")
		os.Exit(1)
	}
	child0 := ctx.GetChild(0).(antlr.ParseTree)
	varNode := v.visitExpr(child0)
	child1 := ctx.GetChild(1)
	op := ""
	if term, ok := child1.(antlr.TerminalNode); ok {
		op = term.GetText()
	} else {
		line := ctx.GetStart().GetLine()
		fmt.Printf("Error on line %d: Expected an operator (e.g., '=', '+=', etc.) but found something else.\n", line)
		os.Exit(1)
	}
	child2 := ctx.GetChild(2).(antlr.ParseTree)
	exprNode := v.visitExpr(child2)
	return &ast.AssignStatementNode{
		Variable:   varNode,
		Operator:   op,
		Expression: exprNode,
	}
}

func (v *CustomVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	child := ctx.GetChild(0).(antlr.ParseTree)
	expr := v.visitExpr(child)

	llvmCode := ""
	var finalExpr ast.Expression

	if lit, ok := expr.(*ast.LiteralNode); ok {
		if line, ok := lit.Value.(string); ok {
			if strings.HasPrefix(line, "local") {
				parts := strings.Split(line, "=")
				if len(parts) == 2 {
					variableName := strings.TrimSpace(strings.ReplaceAll(parts[0], "local", ""))
					initialValue := strings.TrimSpace(parts[1])
					llvmCode = fmt.Sprintf("%%%s = alloca i32\nstore i32 %s, i32* %%%s\n", variableName, initialValue, variableName)
					fmt.Println(llvmCode)
					finalExpr = &ast.VariableNode{Name: variableName}
				}
			} else {
				var operator string
				if strings.Contains(line, "+") {
					operator = "+"
				} else if strings.Contains(line, "-") {
					operator = "*"
				} else if strings.Contains(line, "/") {
					operator = "*"
				} else if strings.Contains(line, "^") {
					operator = "^"
				}

				if operator != "" {
					parts := strings.Split(line, operator)
					if len(parts) == 2 {
						left := &ast.VariableNode{Name: strings.TrimSpace(parts[0])}
						right := &ast.VariableNode{Name: strings.TrimSpace(parts[1])}
						finalExpr = &ast.BinaryExpressionNode{
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
			fmt.Println("Error: Literal value is not a string")
			finalExpr = expr
		}
	} else {
		finalExpr = expr
	}
	
	if varNode, ok := finalExpr.(*ast.VariableNode); ok {
		llvmCode = fmt.Sprintf("load i32, i32* %%%s", varNode.Name)
		fmt.Println(llvmCode)
		return nil
	}

	if binExpr, ok := finalExpr.(*ast.BinaryExpressionNode); ok {
		_ = binExpr
		fmt.Println("Binary Expression")
	}

	return &ast.ExpressionStatementNode{Expr: expr}
}

// VisitReturnStmt corresponds to the returnStatement rule.
// Grammar: returnStatement : RETURN (expression (',' expression)* | functionCall)? ;
func (v *CustomVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	exprs := []ast.Expression{}
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
	return &ast.ReturnStatementNode{Values: exprs}
}

// VisitIfStatement corresponds to the if-statement rule.
func (v *CustomVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	cond := v.visitExpr(ctx.Expression(0).(antlr.ParseTree))
	blocks := ctx.AllBlock()
	thenBlock := &ast.BlockNode{Statements: v.visitBlock(blocks[0])}
	var elseBlock *ast.BlockNode
	if len(blocks) > 1 {
		elseBlock = &ast.BlockNode{Statements: v.visitBlock(blocks[len(blocks)-1])}
	}
	return &ast.IfStatementNode{
		Condition: cond,
		Then:      thenBlock,
		Else:      elseBlock,
	}
}

// Helper: visitBlock collects all statements from a block.
func (v *CustomVisitor) visitBlock(ctx parser.IBlockContext) []ast.Statement {
	stmts := []ast.Statement{}
	for _, sCtx := range ctx.AllStatement() {
		if pt, ok := sCtx.(antlr.ParseTree); ok {
			res := v.Visit(pt)
			if res == nil {
				continue
			}
			if s, ok := res.(ast.Statement); ok {
				stmts = append(stmts, s)
			} else if expr, ok := res.(ast.Expression); ok {
				stmts = append(stmts, &ast.ExpressionStatementNode{Expr: expr})
			}
		}
	}
	return stmts
}
// VisitLiteral corresponds to the literal rule.
func (v *CustomVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
    text := ctx.GetText()
    
    // Handle integers
    if val, err := strconv.Atoi(text); err == nil {
        return &ast.LiteralNode{
            Value: val,
            Type:  "int",
        }
    }
    
    // Handle floating point numbers
    if val, err := strconv.ParseFloat(text, 64); err == nil {
        return &ast.LiteralNode{
            Value: val,
            Type:  "float",
        }
    }
    
    // Handle strings
    if strings.HasPrefix(text, "\"") || strings.HasPrefix(text, "'") {
        return &ast.LiteralNode{
            Value: strings.Trim(text, "\"'"),
            Type:  "string",
        }
    }
    
    // Handle tables or other types
    return &ast.LiteralNode{
        Value: text,
            Type:  "unknown",
        }
    }

// VisitVariable corresponds to the variable rule.
func (v *CustomVisitor) VisitVariable(ctx *parser.VariableContext) interface{} {
	return &ast.VariableNode{Name: ctx.GetText()}
}

// VisitUnaryOperation corresponds to unaryOperation.
func (v *CustomVisitor) VisitUnaryOperation(ctx *parser.UnaryOperationContext) interface{} {
	op := ctx.UnaryOp().GetText()
	expr := v.visitExpr(ctx.Expression().(antlr.ParseTree))
	return &ast.UnaryExpressionNode{
		Operator: op,
		Expr:     expr,
	}
}

// Fallback Visit method.
func (v *CustomVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}