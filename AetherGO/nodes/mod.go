package nodes

import (
	"FLUX/AetherGO/types"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"FLUX/AetherGO/core"
)

// node intefaces that are implemented
var (
	_ types.Statement = (*FunctionDecl)(nil)
	_ types.Statement = (*IfStatement)(nil) 
	_ types.Expression = (*BinaryExpression)(nil)
	_ types.Expression = (*Literal)(nil)
)

// common node types
type (
	Statement = types.Statement
	Expression = types.Expression
	NodeType = types.NodeType
)
