package llvm

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type LLVMBuilder struct {
	Builder     *ir.Block
	curFunc      *ir.Func
	blocks       []*ir.Block
	symbolTable  map[string]value.Value
	variables    map[string]*ir.InstAlloca
	Module       *ir.Module
}

func NewBuilder(module *ir.Module) *LLVMBuilder {
	return &LLVMBuilder{
		Builder:     nil,
		symbolTable: make(map[string]value.Value),
		variables:   make(map[string]*ir.InstAlloca),
		Module:      module,
		blocks:      make([]*ir.Block, 0),
	}
}

func (b *LLVMBuilder) CreateEntryBlock() *ir.Block {
	fn := b.Module.NewFunc("main", types.I32)
	block := fn.NewBlock("entry")
	b.blocks = append(b.blocks, block)
	b.Builder = block
	return block
}

func (b *LLVMBuilder) CurrentBlock() *ir.Block {
	return b.Builder
}

func (b *LLVMBuilder) AddInstruction(inst ir.Instruction) {
	if b.Builder != nil {
		b.Builder.Insts = append(b.Builder.Insts, inst)
	}
}

func (b *LLVMBuilder) CreateAlloca(t types.Type, name string) *ir.InstAlloca {
	return b.CurrentBlock().NewAlloca(t)
}

func (b *LLVMBuilder) GetVariable(name string) value.Value {
	return b.symbolTable[name]
}

func (b *LLVMBuilder) SetVariable(name string, val value.Value) {
	b.symbolTable[name] = val
}
