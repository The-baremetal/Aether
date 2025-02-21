package types

import (
	"encoding/json"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"strings"
)

type (
	Builder = *ir.Block
	Module = ir.Module
)

type Compiler interface {
	GetBuilder() Builder
	GetModule() *Module
	SetVariable(name string, val value.Value)
}

type Statement interface {
	GenIR(c Compiler) error
}

type Expression interface {
	GenIR(c Compiler) (value.Value, error)
}

type Function struct {
	Name        string
	Parameters  []*Variable
	ReturnType  types.Type
	Body        []Statement
	LocalVars   map[string]*Variable
}

type ASTNode struct {
	Type        NodeType
	Value       interface{}
	Children    []*ASTNode
	Annotations map[string]interface{}
	IsJSON      bool
	RawJSON     json.RawMessage
}

type NodeType int

const (
	NodeExpr NodeType = iota
	NodeStmt
	NodeDecl
	NodeLiteral
	NodeControlFlow
	NodeBinaryExpr
	NodePrimaryExpr
	NodeOperator
	NodeProgram
	NodeLocalDecl
	NodeVariable
	NodeFuncCall
	NodeArguments
)

type Variable struct {
	Name         string
	Type         types.Type
	Value        value.Value
	IsConstant   bool
	IsExported   bool
	Metadata     map[string]interface{}
}

type TypeSystem struct {
	types          map[string]types.Type
	currentScope   *SymbolScope
}

func NewTypeSystem() *TypeSystem {
	return &TypeSystem{
		types:        make(map[string]types.Type),
		currentScope: NewSymbolScope(nil),
	}
}

func (ts *TypeSystem) AddType(name string, t types.Type) {
	ts.types[name] = t
}

func (ts *TypeSystem) ResolveType(name string) types.Type {
	if t, ok := ts.types[name]; ok {
		return t
	}
	return nil
}

var NodeTypeStrings = map[NodeType]string{
	NodeExpr:        "Expr",
	NodeStmt:        "Stmt",
	NodeDecl:        "Decl",
	NodeLiteral:     "Literal",
	NodeControlFlow: "ControlFlow",
	NodeBinaryExpr:  "BinaryExpr",
	NodePrimaryExpr: "PrimaryExpr",
	NodeOperator:    "Operator",
	NodeProgram:     "Program",
	NodeLocalDecl:   "LocalDecl",
	NodeVariable:    "Variable",
	NodeFuncCall:    "FuncCall",
	NodeArguments:   "Arguments",
}

func NodeTypeToString(nt NodeType) string {
	if s, ok := NodeTypeStrings[nt]; ok {
		return s
	}
	return "Unknown"
}

func NodeTypeFromString(s string) NodeType {
	switch strings.ToLower(s) {
	case "controlflow": return NodeControlFlow
	case "expr": return NodeExpr
	case "primaryexpr": return NodePrimaryExpr
	case "decl": return NodeDecl
	case "literal": return NodeLiteral
	case "program": return NodeProgram
	case "localdecl": return NodeLocalDecl
	default: return NodeStmt
	}
}