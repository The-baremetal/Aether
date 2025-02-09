package core

import (
    "github.com/antlr4-go/antlr/v4"
    "FLUX/parser"
    "strconv"
    LLVMA "FLUX/AetherGO/LLVMA"
)

type CustomVisitor struct {
	parser.BaseLua_grammar_antlr4Visitor
	Parser *parser.Lua_grammar_antlr4Parser
}

func (v *CustomVisitor) Visit(tree antlr.ParseTree) interface{} {
    return tree.Accept(v)
}

func (v *CustomVisitor) visitNode(node antlr.Tree) interface{} {
    switch ctx := node.(type) {
    case *antlr.TerminalNodeImpl:
        return ctx.GetText()
    case antlr.RuleContext:
        return v.VisitChildren(ctx)
    }
    return nil
}

func (v *CustomVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
    if ctx.GetChildCount() == 1 {
        return v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    }
    if ctx.GetChildCount() == 3 {
        left := v.Visit(ctx.GetChild(0).(antlr.ParseTree)).(LLVMA.Expr)
        right := v.Visit(ctx.GetChild(2).(antlr.ParseTree)).(LLVMA.Expr)
        op := ctx.GetChild(1).GetPayload().(*antlr.CommonToken).GetText()
        
        return &LLVMA.BinaryExpr{
            Left:     left,
            Right:    right,
            Operator: op,
        }
    }
    return nil
}

func (v *CustomVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
    if lit := ctx.Literal(); lit != nil {
        if num := lit.NUMBER(); num != nil {
            value, _ := strconv.Atoi(num.GetText())
            return &LLVMA.IntegerLiteral{Value: value}
        }
    }
    return v.VisitChildren(ctx)
	}

func (v *CustomVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
    var exprs []LLVMA.Expr
    for _, stmt := range ctx.AllStatement() {
        if res := v.Visit(stmt); res != nil {
            if expr, ok := res.(LLVMA.Expr); ok {
                exprs = append(exprs, expr)
            }
        }
    }
    return exprs
}

func (v *CustomVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
    if assign := ctx.AssignStatement(); assign != nil {
        return v.Visit(assign)
    }
    return v.VisitChildren(ctx)
}

func NewCustomVisitor(p *parser.Lua_grammar_antlr4Parser) *CustomVisitor {
	return &CustomVisitor{
		BaseLua_grammar_antlr4Visitor: parser.BaseLua_grammar_antlr4Visitor{},
		Parser: p,
	}
}