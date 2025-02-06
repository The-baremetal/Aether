package resolver

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Resolver struct {
	module       *ir.Module
	currentFunc  *ir.Func
	currentBlock *ir.Block
}

func NewResolver(module *ir.Module, currentFunc *ir.Func, currentBlock *ir.Block) *Resolver {
	return &Resolver{
		module:       module,
		currentFunc:  currentFunc,
		currentBlock: currentBlock,
	}
}

func (r *Resolver) Visit(node interface{}) value.Value {
	switch n := node.(type) {
	case map[string]interface{}:
		rule, _ := n["rule"].(string)
		nodes, _ := n["nodes"].([]interface{})

		switch rule {
		case "functionCall":
			return r.handleFunctionCall(nodes)
		default:
			for _, child := range nodes {
				r.Visit(child)
			}
		}
	case []interface{}:
		for _, child := range n {
			r.Visit(child)
		}
	}
	return nil
}

func (r *Resolver) handleFunctionCall(nodes []interface{}) value.Value {
	if len(nodes) < 1 {
		return nil
	}
	fnName, _ := nodes[0].(string)
	switch fnName {
	case "require":
		return r.require(nodes[1:])
	case "coroutine":
		return r.coroutine(nodes[1:])
	}
	return nil
}

func (r *Resolver) require(args []interface{}) value.Value {
	// (llvm require call (just return the path (eg. local (require as) = require("<path>") just tell llvm to request a link to this during compilation eg. declare i32 @printf(i8*, ...) )))
	return nil
}

func (r *Resolver) coroutine(args []interface{}) value.Value {
	// (llvm coroutine call (just return the path (eg. local (coroutine as) = coroutine("<path>") just tell llvm to request a link to this during compilation eg. i forgot... just import another package in go that handles coroutines because of its complexity )))
	return nil
}

func (r *Resolver) getOrDeclarePuts() *ir.Func {
	if puts := r.module.Funcs["puts"]; puts != nil {
		return puts
	}
	
	return r.module.NewFunc(
		"puts",
		types.I32,
		ir.NewParam("s", types.I8Ptr),
	)
}

func (r *Resolver) createGlobalString(str string) value.Value {
	formattedStr := str + "\x00"
	
	globalName := ".str"
	global := r.module.NewGlobalDef(globalName, constant.NewCharArrayFromString(formattedStr))
	global.Immutable = true
	
	return constant.NewGetElementPtr(
		global.ContentType,
		global,
		constant.NewInt(types.I32, 0),
		constant.NewInt(types.I32, 0),
	)
}
