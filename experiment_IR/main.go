package main

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	// Create a new context
	module := ir.NewModule()
	
	// Create a function with no arguments, returning an integer
	mainFunc := module.NewFunc("main", types.I32)
	
	// Add a basic block
	entry := mainFunc.NewBlock("entry")
	
	// Add some basic LLVM IR operations
	// Create constants
	const1 := constant.NewInt(types.I32, 42)
	const2 := constant.NewInt(types.I32, 58)
	
	// Create an addition instruction
	result := entry.NewAdd(const1, const2)
	
	// Create a return instruction
	entry.NewRet(result)
	
	// Print the module as LLVM IR
	fmt.Println(module)
}
