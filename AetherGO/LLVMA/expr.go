package LLVMA

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
	"FLUX/parser"
	"strconv"
	"fmt"
)

type Expr interface {
	Generate(builder *ir.Block) value.Value
}

type BinaryExpr struct {
	Left     Expr
	Right    Expr
	Operator string
}

type UnaryExpr struct {
	Operand  Expr
	Operator string
}

type IntegerLiteral struct {
	Value int
}

type BooleanLiteral struct {
	Value bool
}

type StringLiteral struct {
	Value string
}

type VariableExpr struct {
	Name string
}

type AssignExpr struct {
	Target Expr
	Value  Expr
}

type ExprGenerator struct {
	Module  *ir.Module
	Builder *ir.Block
}

type ValueWrapper struct {
	Value value.Value
}

func (v *ValueWrapper) Generate(builder *ir.Block) value.Value {
	return v.Value
}

func (e *BinaryExpr) Generate(builder *ir.Block) value.Value {
	left := e.Left.Generate(builder)
	right := e.Right.Generate(builder)
	if left.Type() != right.Type() {
		left, right = convertOperands(builder, left, right)
	}

	switch e.Operator {
	case "+", "..":
		return builder.NewAdd(left, right)
	case "-":
		return builder.NewSub(left, right)
	case "*":
		return builder.NewMul(left, right)
	case "/", "//":
		return builder.NewSDiv(left, right)
	case "%":
		return builder.NewSRem(left, right)
	case "^":
		// NOT COmPLETE
		return callRuntimeFunc(builder, "pow", []value.Value{left, right})
	case "and", "&":
		return builder.NewAnd(left, right)
	case "or", "|":
		return builder.NewOr(left, right)
	case "<<":
		return builder.NewShl(left, right)
	case ">>":
		return builder.NewLShr(left, right)
	case "~":
		return builder.NewXor(left, right)
	default:
		panic("unsupported binary operator: " + e.Operator)
	}
}

func (e *UnaryExpr) Generate(builder *ir.Block) value.Value {
	operand := e.Operand.Generate(builder)
	fmt.Println("you suck")
	switch e.Operator {
	case "-":
		return builder.NewSub(constant.NewInt(types.I32, 0), operand)
	case "not", "~":
		return builder.NewXor(operand, constant.NewInt(types.I32, -1))
	case "#":
		return callRuntimeFunc(builder, "len", []value.Value{operand})
	default:
		panic("unsupported unary operator: " + e.Operator)
	}
}

func (e *IntegerLiteral) Generate(builder *ir.Block) value.Value {
	return constant.NewInt(types.I32, int64(e.Value))
}

func (e *BooleanLiteral) Generate(builder *ir.Block) value.Value {
	return constant.NewInt(types.I1, boolToInt(e.Value))
}

func (e *StringLiteral) Generate(builder *ir.Block) value.Value {
	return callRuntimeFunc(builder, "createString", []value.Value{
		constant.NewCharArray([]byte(e.Value)),
		constant.NewInt(types.I32, int64(len(e.Value))),
	})
}

func (e *VariableExpr) Generate(builder *ir.Block) value.Value {
	for _, bb := range builder.Parent.Blocks {
		for _, inst := range bb.Insts {
			if alloca, ok := inst.(*ir.InstAlloca); ok && alloca.Name() == e.Name {
				return builder.NewLoad(types.I32, alloca)
			}
		}
	}
	panic("undeclared variable: " + e.Name)
}

func (e *AssignExpr) Generate(builder *ir.Block) value.Value {
	target := e.Target.Generate(builder)
	value := e.Value.Generate(builder)
	builder.NewStore(value, target.(*ir.InstAlloca))
	return value
}

func boolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func convertOperands(builder *ir.Block, a, b value.Value) (value.Value, value.Value) {
	if a.Type().Equal(types.I1) {
		a = builder.NewZExt(a, types.I32)
	}
	if b.Type().Equal(types.I1) {
		b = builder.NewZExt(b, types.I32)
	}
	return a, b
}

func WriteExpr(builder *ir.Block, expr Expr) value.Value {
	return expr.Generate(builder)
}

func callRuntimeFunc(b *ir.Block, name string, args []value.Value) value.Value {
	fn := b.Parent.Parent.NewFunc(name, types.I32, ir.NewParam("arg1", types.I8Ptr))
	return b.NewCall(fn, args...)
}

func NewExprGenerator(mod *ir.Module, entry *ir.Block) *ExprGenerator {
	return &ExprGenerator{
		Module:  mod,
		Builder: entry,
	}
}

func (g *ExprGenerator) GenerateExpression(ctx *parser.ExpressionContext) value.Value {
	if len(ctx.AllExpression()) == 2 {
		left := g.GenerateExpression(ctx.Expression(0).(*parser.ExpressionContext))
		right := g.GenerateExpression(ctx.Expression(1).(*parser.ExpressionContext))
		op := ""
		
		if opGroup := ctx.OperatorGroup(0); opGroup != nil {
			op = opGroup.GetText()
			return (&BinaryExpr{
				Left:     &ValueWrapper{Value: left},
				Right:    &ValueWrapper{Value: right},
				Operator: op,
			}).Generate(g.Builder)
		}
	}
	if primary := ctx.PrimaryExpression(); primary != nil {
		return g.GeneratePrimary(primary.(*parser.PrimaryExpressionContext))
	}
	
	panic("unsupported expression type")
}

func (g *ExprGenerator) GeneratePrimary(ctx *parser.PrimaryExpressionContext) value.Value {
    if lit := ctx.Literal(); lit != nil {
        if num := lit.NUMBER(); num != nil {
            val, _ := strconv.Atoi(num.GetText())
            return (&IntegerLiteral{Value: val}).Generate(g.Builder)
        }
    }
    if varExpr := ctx.Variable(); varExpr != nil {
        return (&VariableExpr{Name: varExpr.GetText()}).Generate(g.Builder)
    }
    if fc := ctx.FunctionCall(); fc != nil {
        return g.GenerateFunctionCall(fc.(*parser.FunctionCallContext))
    }
    panic("unsupported primary expression")
}

func (g *ExprGenerator) HandlePrint(value value.Value) {
    var fmtStr *ir.Global
    for _, global := range g.Module.Globals {
        if global.Name() == "fmt.str" {
            fmtStr = global
            break
        }
    }
    
    if fmtStr == nil {
        fmtStr = g.Module.NewGlobalDef("fmt.str", 
            constant.NewCharArrayFromString("%d\n\x00"))
    }
    formatPtr := g.Builder.NewBitCast(fmtStr, types.I8Ptr)
    g.Builder.NewCall(g.Module.Funcs[0], formatPtr, value)
}

func (g *ExprGenerator) GenerateFunctionCall(ctx *parser.FunctionCallContext) value.Value {
    var fnName string
    if varCtx := ctx.Variable(); varCtx != nil {
        fnName = varCtx.GetText()
    }

    var args []value.Value
    for _, exprCtx := range ctx.AllExpression() {
        args = append(args, g.GenerateExpression(exprCtx.(*parser.ExpressionContext)))
    }

    switch fnName {
    case "print":
        if len(args) == 0 {
            panic("print requires at least one argument")
        }
        g.HandlePrint(args[0])
        return nil
    default:
        var fn *ir.Func
        for _, f := range g.Module.Funcs {
            if f.Name() == fnName {
                fn = f
                break
            }
        }
        if fn != nil {
            return g.Builder.NewCall(fn, args...)
        }
        panic("undefined function: " + fnName)
    }
}