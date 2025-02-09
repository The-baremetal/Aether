package LLVMA

import (
	"fmt"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/llir/llvm/ir"
)



func (g *ExprGenerator) HandlePrint(value value.Value) {
	puts := g.Module.NewFunc("puts", types.I32, 
		ir.NewParam("", types.I8Ptr),
	)
	str := g.Builder.NewAlloca(types.NewArray(16, types.I8))
	g.Builder.NewStore(constant.NewCharArrayFromString(fmt.Sprintf("%d\n", value)), str)
	g.Builder.NewCall(puts, str)
} 