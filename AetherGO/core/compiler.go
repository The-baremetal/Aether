package core

import (
    "FLUX/AetherGO/llvm"
    "FLUX/AetherGO/types"
    "github.com/llir/llvm/ir"
    "github.com/llir/llvm/ir/value"
    irtypes "github.com/llir/llvm/ir/types"
    "github.com/llir/llvm/ir/constant"
    "fmt"
    "encoding/json"
)

type Compiler struct {
    Module      *ir.Module
    Builder     *llvm.LLVMBuilder
    curFunc     *ir.Func
    blocks      []*ir.Block
    Scope       *types.SymbolScope
    variables   map[string]*types.Variable
    TypeSystem  *types.TypeSystem
    IRGenConfig struct {
        BinaryOps map[string]IRGenRule
    }
    DebugInfo   *DebugInfo
}
type (
    IRGenRule struct {
        Op    string
        GenFn func(*Compiler, value.Value, value.Value) (value.Value, error)
    }
    
    SymbolScope struct {
        Parent  *SymbolScope
        Symbols map[string]*types.Variable
    }
    
    DebugInfo struct {
        SourceLines []string
        FilePath    string
    }
)

func NewCompiler() *Compiler {
    module := ir.NewModule()
    return &Compiler{
        Builder:     llvm.NewBuilder(module),
        TypeSystem:  types.NewTypeSystem(),
        Scope:       types.NewSymbolScope(nil),
        Module:      module,
        IRGenConfig: struct {
            BinaryOps map[string]IRGenRule
        }{
            BinaryOps: make(map[string]IRGenRule),
        },
    }
}

func (c *Compiler) Compile(ast *types.ASTNode) error {
    fmt.Println("=== AST DUMP ===")
    fmt.Println(c.DumpAST(ast))
    fmt.Println("================")
    
    c.Module.NewGlobalDef("print_int_fmt", constant.NewCharArrayFromString("%d\n\x00"))
    c.Module.NewGlobalDef("print_str_fmt", constant.NewCharArrayFromString("%s\n\x00"))
    
    printf := c.Module.NewFunc("printf", irtypes.I32, ir.NewParam("format", irtypes.I8Ptr))
    printf.Sig.Variadic = true
    
    c.Builder.CreateEntryBlock()
    
    for _, child := range ast.Children {
        if stmt, ok := child.Value.(Statement); ok {
            if err := stmt.GenIR(c); err != nil {
                return err
            }
        }
    }
    if c.Builder.CurrentBlock().Term == nil {
        c.Builder.CurrentBlock().NewRet(constant.NewInt(irtypes.I32, 0))
    }
    return nil
}

func (c *Compiler) EnterScope() {
    c.Scope = types.NewSymbolScope(c.Scope)
}

func (c *Compiler) ExitScope() {
    if c.Scope.Parent != nil {
        c.Scope = c.Scope.Parent
    }
}

func (c *Compiler) AddIRRule(op string, genFn func(*Compiler, value.Value, value.Value) (value.Value, error)) {
    c.IRGenConfig.BinaryOps[op] = IRGenRule{Op: op, GenFn: genFn}
}

func (c *Compiler) Finalize() string {
    return c.Module.String()
}

func (c *Compiler) DumpAST(ast *types.ASTNode) string {
    fmt.Println("AST")
    return FormatAST(ast, 0)
}

func (c *Compiler) compileFromJSON(raw json.RawMessage) error {
    var jsonAST struct {
        NodeType string
        Value    interface{}
        Children []json.RawMessage
        IsJSON   bool
    }
    
    if err := json.Unmarshal(raw, &jsonAST); err != nil {
        return fmt.Errorf("JSON AST parsing failed: %w", err)
    }
    
    ast := &types.ASTNode{
        Type:   types.NodeTypeFromString(jsonAST.NodeType),
        Value:  jsonAST.Value,
        IsJSON: true,
    }
    
    for _, child := range jsonAST.Children {
        var childNode types.ASTNode
        if err := json.Unmarshal(child, &childNode); err != nil {
            return err
        }
        ast.Children = append(ast.Children, &childNode)
    }
    for _, child := range ast.Children {
        if stmt, ok := child.Value.(Statement); ok {
            if err := stmt.GenIR(c); err != nil {
                return err
            }
        }
    }
    
    return nil
}