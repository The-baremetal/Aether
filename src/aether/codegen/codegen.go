package codegen

import (
	"FLUX/src/aether/ast"
	"fmt"
)

import "sync/atomic"

// Leash represents the ownership of a variable.
type Leash struct {
	owner   string
	mutable atomic.Bool
}

// variableLeashes tracks the leashes for each variable.
var variableLeashes = make(map[string]*Leash)

// acquireLeash acquires the leash for a variable.
func acquireLeash(variable string, owner string, mutable bool) {
	leash, ok := variableLeashes[variable]
	if ok {
		if mutable {
			if leash.mutable.Load() {
				// TODO: Handle case where mutable leash is already acquired.
				return
			}
			leash.mutable.Store(true)
		}
		// TODO: Handle case where leash is already acquired.
		return
	}
	leash = &Leash{owner: owner}
	if mutable {
		leash.mutable.Store(true)
	}
	variableLeashes[variable] = leash
}

// releaseLeash releases the leash for a variable.
func releaseLeash(variable string) {
	delete(variableLeashes, variable)
}

// generateLLVMCodeForBinaryExpression generates LLVM code for a BinaryExpressionNode.
func GenerateLLVMCodeForBinaryExpression(b *ast.BinaryExpressionNode) string {
	left := GenerateLLVMCodeForExpression(b.Left)
	right := GenerateLLVMCodeForExpression(b.Right)
	switch b.Operator {
	case "+":
		return fmt.Sprintf("add i32 %s, %s", left, right)
	case "-":
		return fmt.Sprintf("sub i32 %s, %s", left, right)
	case "*":
		return fmt.Sprintf("mul i32 %s, %s", left, right)
	case "/":
		return fmt.Sprintf("sdiv i32 %s, %s", left, right)
	default:
		return ""
	}
}
func GenerateLLVMCodeForExpression(e ast.Expression) string {
	// TODO: Implement leash system here
	switch v := e.(type) {
	case *ast.LiteralNode:
		if str, ok := v.Value.(string); ok {
			return str
		} else {
			return ""
		}
	case *ast.VariableNode:
		variable := v.Name
		acquireLeash(variable, "GenerateLLVMCodeForExpression", v.Mutable)
		defer releaseLeash(variable)
		return "%" + variable
	case *ast.BinaryExpressionNode:
		return GenerateLLVMCodeForBinaryExpression(v)
	case *ast.UnaryExpressionNode:
		return GenerateLLVMCodeForUnaryOperation(v)
	default:
		return ""
	}
}
func GenerateLLVMCodeForUnaryOperation(u *ast.UnaryExpressionNode) string {
	expr := GenerateLLVMCodeForExpression(u.Expr)
	if u.Operator == "-" {
		return fmt.Sprintf("sub i32 0, %s", expr)
	}
	return ""
}
