package llvm

import (
	"FLUX/AetherGO/types"
	"github.com/llir/llvm/ir"
)

type CompilationContext struct {
	Module      *ir.Module
	Builder     *LLVMBuilder
	TypeSystem  *types.TypeSystem
	CurrentFunc *ir.Func
	SymbolTable *types.SymbolScope
}

func NewContext() *CompilationContext {
	module := ir.NewModule()
	return &CompilationContext{
		Module:      module,
		Builder:     NewBuilder(module),
		TypeSystem:  types.NewTypeSystem(),
		SymbolTable: types.NewSymbolScope(nil),
	}
}

func (ctx *CompilationContext) EnterScope() {
	ctx.SymbolTable = types.NewSymbolScope(ctx.SymbolTable)
}

func (ctx *CompilationContext) ExitScope() {
	if ctx.SymbolTable.Parent != nil {
		ctx.SymbolTable = ctx.SymbolTable.Parent
	}
}

func (ctx *CompilationContext) CurrentBlock() *ir.Block {
	return ctx.Builder.CurrentBlock()
}
