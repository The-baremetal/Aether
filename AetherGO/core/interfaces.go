package core

import (
	"github.com/llir/llvm/ir"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"FLUX/AetherGO/types"
)

// core compiler interfaces
type (
    Statement interface {
        GenIR(c *Compiler) error
    }
    
    Expression interface {
        GenIR(c *Compiler) (value.Value, error)
    }
)

// export common stuff
type (
	Type = irtypes.Type
	Value = value.Value
	Module = ir.Module
)

func (c *Compiler) GetBuilder() *ir.Block { return c.Builder.Builder }
func (c *Compiler) GetModule() *ir.Module { return c.Module }
func (c *Compiler) SetVariable(name string, val value.Value) {
    c.variables[name] = &types.Variable{Value: val}
}
