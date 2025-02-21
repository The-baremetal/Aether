package nodes

import (
    "FLUX/AetherGO/core"
    "github.com/llir/llvm/ir"
    "github.com/llir/llvm/ir/types"
)

type FunctionDecl struct {
    Name       string
    Parameters []*core.Variable
    ReturnType ir.Type
    Body       []interface{}
}

func (fd *FunctionDecl) GenIR(c *core.Compiler) error {
    paramTypes := make([]types.Type, len(fd.Parameters))
    for i, param := range fd.Parameters {
        paramTypes[i] = param.Type
    }
    
    fnType := types.NewFunc(fd.ReturnType, paramTypes...)
    function := c.Builder.Module.NewFunc(fd.Name, fnType)
    entry := function.NewBlock("entry")
    c.Builder.blocks = append(c.Builder.blocks, entry)
    for i, param := range fd.Parameters {
        alloca := entry.NewAlloca(param.Type)
        entry.NewStore(function.Params[i], alloca)
        c.Builder.SetVariable(param.Name, alloca)
    }
    for _, stmt := range fd.Body {
        if err := core.GenerateStatementIR(c, stmt); err != nil {
            return err
        }
    }
    if fd.ReturnType == types.Void {
        entry.NewRet(nil)
    } else {
        entry.NewRet(ir.NewInt(fd.ReturnType, 0))
    }
    
    return nil
}
