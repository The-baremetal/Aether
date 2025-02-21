package llvm

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/llir/llvm/ir/constant"
)

func CreateConstant(t types.Type, val interface{}) constant.Constant {
	switch typ := t.(type) {
	case *types.IntType:
		return constant.NewInt(typ, val.(int64))
	case *types.FloatType:
		return constant.NewFloat(typ, val.(float64))
	case *types.PointerType:
		if typ.ElemType == types.I8 {
			return constant.NewCharArrayFromString(val.(string))
		}
		return constant.NewNull(typ)
	default:
		if ptrType, ok := t.(*types.PointerType); ok {
			return constant.NewNull(ptrType)
		}
		return nil
	}
}

func GetZeroValue(t types.Type) value.Value {
	switch typ := t.(type) {
	case *types.IntType:
		return constant.NewInt(typ, 0)
	case *types.PointerType:
		return constant.NewNull(typ)
	default:
		return nil
	}
}

func boolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
