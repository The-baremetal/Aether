package core

import (
	"fmt"
	"FLUX/AetherGO/types"
	llvmtypes "github.com/llir/llvm/ir/types"
)

type TypeChecker struct {
	ts *types.TypeSystem
}

func NewTypeChecker() *TypeChecker {
	return &TypeChecker{
		ts: types.NewTypeSystem(),
	}
}

func (tc *TypeChecker) CheckTypeCompatibility(expected, actual llvmtypes.Type) error {
	if llvmtypes.Equal(expected, actual) {
		return nil
	}
	if isNumericType(expected) && isNumericType(actual) {
		return nil
	}
	if ptrType, ok := expected.(*llvmtypes.PointerType); ok {
		if actual.Equal(llvmtypes.NewPointer(ptrType.ElemType)) {
			return nil
		}
	}
	return fmt.Errorf("type mismatch: expected %s, got %s", expected, actual)
}

func (tc *TypeChecker) CheckBinaryOperation(op string, left, right llvmtypes.Type) (llvmtypes.Type, error) {
	switch op {
	case "+", "-", "*", "/":
		if isNumericType(left) && isNumericType(right) {
			return widestType(left, right), nil
		}
		return nil, fmt.Errorf("operator %s requires numeric operands", op)
	case "==", "!=", "<", ">":
		if llvmtypes.Equal(left, right) {
			return llvmtypes.I1, nil
		}
		return nil, fmt.Errorf("comparison operands must be same type")
	default:
		return nil, fmt.Errorf("unsupported operator: %s", op)
	}
}

func (tc *TypeChecker) CheckFunctionSignature(name string, params []llvmtypes.Type, returnType llvmtypes.Type) error {
	fnType := tc.ts.ResolveType(name)
	if fnType == nil {
		return fmt.Errorf("undefined function %s", name)
	}
	
	sig, ok := fnType.(*llvmtypes.FuncType)
	if !ok {
		return fmt.Errorf("%s is not a function", name)
	}
	
	if len(params) != len(sig.Params) {
		return fmt.Errorf("argument count mismatch for %s", name)
	}
	
	for i, param := range sig.Params {
		if err := tc.CheckTypeCompatibility(param, params[i]); err != nil {
			return fmt.Errorf("parameter %d: %w", i+1, err)
		}
	}
	
	if !llvmtypes.Equal(returnType, sig.RetType) {
		return fmt.Errorf("return type mismatch for %s", name)
	}
	return nil
}

func isNumericType(t llvmtypes.Type) bool {
	return t.Equal(llvmtypes.I32) || t.Equal(llvmtypes.I64) || t.Equal(llvmtypes.Float)
}

func widestType(a, b llvmtypes.Type) llvmtypes.Type {
	typeOrder := map[llvmtypes.Type]int{
		llvmtypes.I32:  1,
		llvmtypes.I64:  2,
		llvmtypes.Float: 3,
	}
	if typeOrder[a] > typeOrder[b] {
		return a
	}
	return b
}