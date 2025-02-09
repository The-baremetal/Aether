package core

import (
    "fmt"
    "github.com/antlr4-go/antlr/v4"
    "FLUX/parser"
    "github.com/llir/llvm/ir"
    LLVMA "FLUX/AetherGO/LLVMA"
    "github.com/llir/llvm/ir/constant"
    "github.com/llir/llvm/ir/types"
    "github.com/llir/llvm/ir/value"
)

type Variable struct {
    Name string
    Type types.Type
    Value value.Value
}

type Compiler struct {
    parser    *parser.Lua_grammar_antlr4Parser
    visitor   *CustomVisitor
    module    *ir.Module
    mainFn    *ir.Func
    entry     *ir.Block
    variables map[string]*Variable
    exprGen   *LLVMA.ExprGenerator
}

func NewCompiler() *Compiler {
    return &Compiler{}
}

func (c *Compiler) SetupParser(input string) error {
    is := antlr.NewInputStream(input)
    lexer := parser.NewLua_grammar_antlr4Lexer(is)
    stream := antlr.NewCommonTokenStream(lexer, 0)
    c.parser = parser.NewLua_grammar_antlr4Parser(stream)
    tree := c.parser.Program()
    fmt.Println("Raw AST structure:")
    fmt.Println(antlr.TreesStringTree(tree, c.parser.GetRuleNames(), c.parser))
    
    c.visitor = NewCustomVisitor(c.parser)
    return nil
}

func (c *Compiler) GenerateIR() error {
    c.variables = make(map[string]*Variable)
    
    var err error
    if c.module, err = LLVMA.InitModule(); err != nil {
        return fmt.Errorf("module init failed: %w", err)
    }

    printf := c.module.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
    printf.Sig.Variadic = true
    
    if c.mainFn, err = LLVMA.InitMain(c.module); err != nil {
        return fmt.Errorf("main function init failed: %w", err)
    }
    if c.entry, err = LLVMA.InitBlock(c.mainFn); err != nil {
        return fmt.Errorf("entry block init failed: %w", err)
    }
    c.exprGen = LLVMA.NewExprGenerator(c.module, c.entry)
    
    c.entry.NewRet(constant.NewInt(types.I32, 0))
    return nil
}

func (c *Compiler) ProcessAST() error {
    tree := c.parser.Program()
    ctx := tree.(*parser.ProgramContext)
    
    for _, stmt := range ctx.AllStatement() {
        if assign := stmt.AssignStatement(); assign != nil {
            c.handleAssignment(assign)
        } else if expr := stmt.Expression(); expr != nil {
            if primaryExpr := expr.PrimaryExpression(); primaryExpr != nil {
                if fc := primaryExpr.FunctionCall(); fc != nil {
                    c.handleFunctionCall(fc)
                }
            }
        } else if localDecl := stmt.LocalDeclaration(); localDecl != nil {
            c.handleLocalDeclaration(localDecl.(*parser.LocalDeclarationContext))
        }
    }
    
    return nil
}

func (c *Compiler) handleLocalDeclaration(ctx *parser.LocalDeclarationContext) {
    ident := ctx.IDENTIFIER(0).GetText()
    alloc := c.entry.NewAlloca(types.I32)
    alloc.SetName(ident)
    
    if expr := ctx.Expression(); expr != nil {
        val := c.visitor.Visit(expr).(value.Value)
        c.entry.NewStore(val, alloc)
    }
    
    c.variables[ident] = &Variable{
        Name: ident,
        Type: types.I32,
        Value: alloc,
    }
}

func (c *Compiler) handleExpression(ctx *parser.ExpressionContext) value.Value {
    return c.exprGen.GenerateExpression(ctx)
}

func (c *Compiler) handleFunctionCall(ctx parser.IFunctionCallContext) {
    var fnName string
    if varCtx := ctx.Variable(); varCtx != nil {
        fnName = varCtx.GetText()
    }

    var args []value.Value
    for _, exprCtx := range ctx.AllExpression() {
        args = append(args, c.exprGen.GenerateExpression(exprCtx.(*parser.ExpressionContext)))
    }

    switch fnName {
    case "print":
        c.exprGen.HandlePrint(args[0])
    default:
        var fn *ir.Func
        for _, f := range c.module.Funcs {
            if f.Name() == fnName {
                fn = f
                break
            }
        }
        if fn != nil {
            c.entry.NewCall(fn, args...)
        }
    }
}

func (c *Compiler) handleAssignment(ctx parser.IAssignStatementContext) {
    varIdent := ctx.Variable().GetText()
    variable := c.variables[varIdent]
    val := c.exprGen.GenerateExpression(ctx.Expression().(*parser.ExpressionContext))
    c.entry.NewStore(val, variable.Value)
}

func (c *Compiler) Finalize() (string, error) {
    return c.module.String(), nil
}