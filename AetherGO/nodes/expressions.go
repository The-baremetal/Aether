package nodes

import (
    "FLUX/AetherGO/llvm"
    "github.com/llir/llvm/ir"
    "github.com/llir/llvm/ir/types"
    "github.com/llir/llvm/ir/value"
    "github.com/llir/llvm/ir/constant"
    "fmt"
)

type BinaryExpression struct {
    Op    string
    Left  interface{}
    Right interface{}
}

func (b *BinaryExpression) GenIR(c *core.Compiler) (value.Value, error) {
    leftVal, err := core.GenerateExpressionIR(c, b.Left)
    if err != nil {
        return nil, err
    }
    rightVal, err := core.GenerateExpressionIR(c, b.Right)
    if err != nil {
        return nil, err
    }
    
    if rule, exists := c.IRGenConfig.BinaryOps[b.Op]; exists {
        return rule.GenFn(c, leftVal, rightVal)
    }
    return nil, core.NewCompileError("unsupported binary operator: " + b.Op)
}

type Literal struct {
    Type  types.Type
    Value interface{}
}

func (l *Literal) GenIR(c *core.Compiler) (value.Value, error) {
    switch l.Type {
    case types.I32:
        return ir.NewInt(types.I32, l.Value.(int64)), nil
    case types.I1:
        return ir.NewInt(types.I1, boolToInt(l.Value.(bool))), nil
    case types.I8Ptr:
        globalName := fmt.Sprintf("str.%d", c.Module.Globals.Len())
        str := constant.NewCharArrayFromString(l.Value.(string) + "\x00")
        global := c.Module.NewGlobalDef(globalName, str)
        return global, nil
    default:
        return nil, core.NewCompileError("unsupported literal type")
    }
}

func boolToInt(b bool) int64 {
    if b {
        return 1
    }
    return 0
}
