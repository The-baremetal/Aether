package llvm

import (
	"github.com/llir/llvm/ir"
	llvmtypes "github.com/llir/llvm/ir/types"
)

type (
	Module  = ir.Module
	Builder = *ir.Block
	Type    = llvmtypes.Type
)

var (
	NewModule = ir.NewModule
	NewFunc = ir.NewFunc
)