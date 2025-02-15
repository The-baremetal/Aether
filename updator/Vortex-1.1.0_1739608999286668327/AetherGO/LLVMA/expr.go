package LLVMA

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
	"FLUX/parser"
	"strconv"
	"fmt"
	"FLUX/AetherGO/utils"
	"github.com/llir/llvm/ir/enum"
	"strings"
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
	Value    string
	generator *ExprGenerator
}

type VariableStorage interface {
    GetVariableType(name string) types.Type
    GetVariableValue(name string) value.Value
}

type VariableExpr struct {
	Name    string
	Storage VariableStorage
}

type AssignExpr struct {
	Target Expr
	Value  Expr
}

type ExprGenerator struct {
	Module    *ir.Module
	Builder   *ir.Block
	Storage   VariableStorage
	counter   int
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
	case ">":
		return builder.NewICmp(enum.IPredSGT, left, right)
	case "<":
		return builder.NewICmp(enum.IPredSLT, left, right)
	case ">=":
		return builder.NewICmp(enum.IPredSGE, left, right)
	case "<=":
		return builder.NewICmp(enum.IPredSLE, left, right)
	case "==":
		return builder.NewICmp(enum.IPredEQ, left, right)
	case "~=":
		return builder.NewICmp(enum.IPredNE, left, right)
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
		utils.LogError("unsupported binary operator: %s", e.Operator)
		panic("operator error")
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

func (s *StringLiteral) Generate(builder *ir.Block) value.Value {
	globalName := fmt.Sprintf("str.%d", s.generator.counter)
	s.generator.counter++
	arrType := types.NewArray(uint64(len(s.Value)+1), types.I8)
	global := s.generator.Module.NewGlobalDef(globalName, 
		constant.NewCharArrayFromString(s.Value+"\x00"))
	zero := constant.NewInt(types.I32, 0)
	gep := builder.NewGetElementPtr(arrType, global, zero, zero)
	return builder.NewBitCast(gep, types.I8Ptr)
}

func (v *VariableExpr) Generate(builder *ir.Block) value.Value {
	varType := v.Storage.GetVariableType(v.Name)
	varValue := v.Storage.GetVariableValue(v.Name)
	if varType == nil || varValue == nil {
		panic("undefined variable: " + v.Name)
	}
	return builder.NewLoad(varType, varValue)
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

func NewExprGenerator(mod *ir.Module, entry *ir.Block, storage VariableStorage) *ExprGenerator {
	return &ExprGenerator{
		Module:  mod,
		Builder: entry,
		Storage: storage,
	}
}
var ExprRegistry map[string]func(*ExprGenerator, parser.IExpressionContext) value.Value
func (g *ExprGenerator) GenerateExpression(ctx parser.IExpressionContext) value.Value {
	exprType := getExprType(ctx)
	if handler, exists := ExprRegistry[exprType]; exists {
		return handler(g, ctx)
	}
	panic("Unsupported expression type: " + exprType)
}
func handlePrimary(g *ExprGenerator, ctx parser.IExpressionContext) value.Value {
	prim := ctx.PrimaryExpression().(*parser.PrimaryExpressionContext)
	return g.GeneratePrimary(prim)
}
func init() {
	ExprRegistry = map[string]func(*ExprGenerator, parser.IExpressionContext) value.Value{
		"Binary":  handleBinary,
		"Unary":   handleUnary,
		"Call":    handleCall,
		"Primary": handlePrimary,
	}
}

func handleBinary(g *ExprGenerator, ctx parser.IExpressionContext) value.Value {
	if len(ctx.AllOperatorGroup()) == 0 {
		panic("binary expression without operator")
	}
	
	var val value.Value
	if prim := ctx.PrimaryExpression(); prim != nil {
		val = g.GeneratePrimary(prim.(*parser.PrimaryExpressionContext))
	} else {
		panic("binary expression missing left operand")
	}
	
	for i := 0; i < len(ctx.AllOperatorGroup()); i++ {
		op := ctx.OperatorGroup(i).GetText()
		right := g.GenerateExpression(ctx.Expression(i))
		
		switch op {
		case "+", "..":
			val = g.Builder.NewAdd(val, right)
		case "-":
			val = g.Builder.NewSub(val, right)
		case "*":
			val = g.Builder.NewMul(val, right)
		case "/":
			val = g.Builder.NewSDiv(val, right)
		case ">", "<", ">=", "<=", "==", "~=":
			val = g.Builder.NewICmp(getComparisonPredicate(op), val, right)
		default:
			panic("unsupported binary operator: " + op)
		}
	}
	return val
}

func handleUnary(g *ExprGenerator, ctx parser.IExpressionContext) value.Value {
	if ctx.UnaryOp() == nil {
		panic("unary expression without operator")
	}
	
	op := ctx.UnaryOp().GetText()
	operand := g.GenerateExpression(ctx.Expression(0))
	
	switch op {
	case "-":
		return g.Builder.NewSub(constant.NewInt(types.I32, 0), operand)
	case "not":
		return g.Builder.NewICmp(enum.IPredEQ, operand, constant.NewInt(types.I32, 0))
	default:
		panic("unsupported unary operator: " + op)
	}
}

func handleCall(g *ExprGenerator, ctx parser.IExpressionContext) value.Value {
	fc := ctx.PrimaryExpression().(*parser.PrimaryExpressionContext).FunctionCall()
	if fc == nil {
		panic("function call context missing")
	}
	return g.GenerateFunctionCall(fc.(*parser.FunctionCallContext))
}

func getExprType(ctx parser.IExpressionContext) string {
	switch {
	case len(ctx.AllOperatorGroup()) > 0:
		return "Binary"
	case ctx.UnaryOp() != nil:
		return "Unary"
	case ctx.PrimaryExpression() != nil:
		return "Primary"
	default:
		return "Unknown"
	}
}

func (g *ExprGenerator) GeneratePrimary(ctx *parser.PrimaryExpressionContext) value.Value {
	if expr := ctx.Expression(); expr != nil {
		return g.GenerateExpression(expr.(parser.IExpressionContext))
	}
	if lit := ctx.Literal(); lit != nil {
		if num := lit.NUMBER(); num != nil {
			val, _ := strconv.Atoi(num.GetText())
			return (&IntegerLiteral{Value: val}).Generate(g.Builder)
		}
		if str := lit.STRING(); str != nil {
			strVal := strings.Trim(str.GetText(), `"`)
			return (&StringLiteral{
				Value:    strVal,
				generator: g,
			}).Generate(g.Builder)
		}
	}
	if varExpr := ctx.Variable(); varExpr != nil {
		return (&VariableExpr{Name: varExpr.GetText(), Storage: g.Storage}).Generate(g.Builder)
	}
	if fc := ctx.FunctionCall(); fc != nil {
		return g.GenerateFunctionCall(fc.(*parser.FunctionCallContext))
	}
	panic("unsupported primary expression")
}

func (g *ExprGenerator) HandlePrint(value value.Value) {
	var formatStr string
	switch value.Type().(type) {
	case *types.PointerType:
		formatStr = "%s\n\x00"
	default:
		formatStr = "%d\n\x00"
	}
	
	fmtName := fmt.Sprintf("fmt.str.%d", g.counter)
	g.counter++
	fmtGlobal := g.Module.NewGlobalDef(fmtName, constant.NewCharArrayFromString(formatStr))
	formatPtr := g.Builder.NewBitCast(fmtGlobal, types.I8Ptr)
	g.Builder.NewCall(g.Module.Funcs[0], formatPtr, value)
}

func (g *ExprGenerator) GenerateFunctionCall(ctx *parser.FunctionCallContext) value.Value {
	var fnName string
	if varCtx := ctx.Variable(); varCtx != nil {
		fnName = varCtx.GetText()
	}

	var args []value.Value
	for _, exprCtx := range ctx.AllExpression() {
		args = append(args, g.GenerateExpression(exprCtx.(parser.IExpressionContext)))
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

func (g *ExprGenerator) GenerateUnary(ctx *parser.ExpressionContext) value.Value {
	unaryOp := ctx.UnaryOp().GetText()
	operandExpr := ctx.Expression(0).(*parser.ExpressionContext)
	operand := g.GenerateExpression(operandExpr)
	
	return (&UnaryExpr{
		Operand:  &ValueWrapper{Value: operand},
		Operator: unaryOp,
	}).Generate(g.Builder)
}

func (g *ExprGenerator) GenerateBinaryOp(left, right value.Value, op string) value.Value {
	return (&BinaryExpr{
		Left:     &ValueWrapper{Value: left},
		Right:    &ValueWrapper{Value: right},
		Operator: op,
	}).Generate(g.Builder)
}

func ValidateExprRegistry() error {
	requiredHandlers := []string{"Binary", "Unary", "Call", "Primary"}
	for _, key := range requiredHandlers {
		if _, exists := ExprRegistry[key]; !exists {
			return fmt.Errorf("missing required expression handler: %s", key)
		}
	}
	return nil
}

var uniqueIDCounter int

func getUniqueID() int {
	uniqueIDCounter++
	return uniqueIDCounter
}

func (g *ExprGenerator) GetUniqueID() int {
	g.counter++
	return g.counter
}

func getComparisonPredicate(op string) enum.IPred {
	switch op {
	case ">": return enum.IPredSGT
	case "<": return enum.IPredSLT 
	case ">=": return enum.IPredSGE
	case "<=": return enum.IPredSLE
	case "==": return enum.IPredEQ
	case "~=": return enum.IPredNE
	default: panic("unsupported comparison operator: " + op)
	}
}