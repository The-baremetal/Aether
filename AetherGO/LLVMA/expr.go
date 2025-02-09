package LLVMA

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
	"FLUX/parser"
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
	for _, inst := range builder.Parent.Blocks[0].Insts {
		if alloca, ok := inst.(*ir.InstAlloca); ok && alloca.Name() == e.Name {
			return builder.NewLoad(types.I32, alloca)
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
	if primary := ctx.PrimaryExpression(); primary != nil {
		return g.GeneratePrimary(primary.(*parser.PrimaryExpressionContext))
	}
	
	if len(ctx.AllExpression()) == 2 {
		left := &IntegerLiteral{Value: 0}
		right := &IntegerLiteral{Value: 0}
		op := ""
		
		if opGroup := ctx.OperatorGroup(0); opGroup != nil {
			op = opGroup.GetText()
			return (&BinaryExpr{
				Left:     left,
				Right:    right,
				Operator: op,
			}).Generate(g.Builder)
		}
	}
	
	panic("unsupported expression type")
}

func (g *ExprGenerator) GeneratePrimary(ctx *parser.PrimaryExpressionContext) value.Value {
	if lit := ctx.Literal(); lit != nil {
		return (&IntegerLiteral{Value: 0}).Generate(g.Builder)
	}
	if varExpr := ctx.Variable(); varExpr != nil {
		return (&VariableExpr{Name: "temp"}).Generate(g.Builder)
	}
	panic("unsupported primary expression")
}