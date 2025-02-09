// really not that hard to understand, im just following llvm documentation
package LLVMA

import (
	LLVM "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"fmt"
)
func InitModule() (*LLVM.Module, error) {
	mod := LLVM.NewModule()
	if mod == nil {
		return nil, fmt.Errorf("failed to create new LLVM module")
	}
	return mod, nil
}
func InitMain(m *LLVM.Module) (*LLVM.Func, error) {
	if m == nil {
		return nil, fmt.Errorf("module is nil, cannot create main function")
	}
	mainFunc := m.NewFunc("main", types.I32)
	if mainFunc == nil {
		return nil, fmt.Errorf("failed to create main function")
	}
	return mainFunc, nil
}
func InitBlock(main *LLVM.Func) (*LLVM.Block, error) {
	if main == nil {
		return nil, fmt.Errorf("main function is nil, cannot create entry block")
	}
	entryBlock := main.NewBlock("entry")
	if entryBlock == nil {
		return nil, fmt.Errorf("failed to create entry block")
	}
	return entryBlock, nil
}
