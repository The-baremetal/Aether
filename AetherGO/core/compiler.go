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
    printf    *ir.Func
    ast       antlr.Tree
}

var counter int

func (c *Compiler) getUniqueID() int {
    counter++
    return counter
}

func NewCompiler() *Compiler {
    return &Compiler{}
}

func (c *Compiler) SetupParser(input string) error {
    fmt.Println("\n=== Setting up parser ===")
    is := antlr.NewInputStream(input)
    lexer := parser.NewLua_grammar_antlr4Lexer(is)
    stream := antlr.NewCommonTokenStream(lexer, 0)
    c.parser = parser.NewLua_grammar_antlr4Parser(stream)
    tree := c.parser.Program()
    c.ast = tree
    fmt.Println("Raw AST structure:")
    fmt.Println(antlr.TreesStringTree(c.ast, c.parser.GetRuleNames(), c.parser))
    
    c.visitor = NewCustomVisitor(c.parser)
    fmt.Println("AST structure printed above")
    return nil
}

func (c *Compiler) GenerateIR() error {
    fmt.Println("\n=== Generating IR ===")
    fmt.Println("Initializing module...")
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
    
    c.printf = printf
    
    fmt.Printf("Created main function: %s\n", c.mainFn.Name())
    fmt.Printf("Created entry block: %s\n", c.entry.Name())
    fmt.Println("Initialized expression generator")
    
    return nil
}

func (c *Compiler) ProcessAST() error {
    fmt.Println("\n=== Processing AST ===")
    fmt.Println("Full AST structure:")
    fmt.Println(antlr.TreesStringTree(c.ast, nil, nil))
    
    ctx := c.ast.(*parser.ProgramContext)
    stmtList := ctx.AllStatement()
    fmt.Printf("Found %d statements\n", len(stmtList))
    
    for i, stmt := range stmtList {
        fmt.Printf("\nProcessing statement %d (%T)\n", i, stmt)
        if assign := stmt.AssignStatement(); assign != nil {
            c.handleAssignment(assign)
        } else if expr := stmt.Expression(); expr != nil {
            c.handleExpression(expr.(antlr.ParserRuleContext))
            if primaryExpr := expr.PrimaryExpression(); primaryExpr != nil {
                if fc := primaryExpr.FunctionCall(); fc != nil {
                    fmt.Println("here is being called")
                    c.handleFunctionCall(fc.(parser.IFunctionCallContext))
                }

            }
        } else if localDecl := stmt.LocalDeclaration(); localDecl != nil {
            c.handleLocalDeclaration(localDecl.(antlr.ParserRuleContext))
        }
    }
    
    return nil
}

func (c *Compiler) handleLocalDeclaration(ctx antlr.ParserRuleContext) {
    localCtx := ctx.(*parser.LocalDeclarationContext)
    ident := localCtx.IDENTIFIER(0).GetText()
    alloc := c.entry.NewAlloca(types.I32)
    alloc.SetName(ident)
    
    if expr := localCtx.Expression(); expr != nil {
        val := c.exprGen.GenerateExpression(expr.(*parser.ExpressionContext))
        c.entry.NewStore(val, alloc)
    }
    
    c.variables[ident] = &Variable{
        Name: ident,
        Type: types.I32,
        Value: alloc,
    }
}

func (c *Compiler) handleExpression(ctx antlr.ParserRuleContext) value.Value {
    fmt.Println("Handling expression")
    exprCtx := ctx.(*parser.ExpressionContext)
    val := c.exprGen.GenerateExpression(exprCtx)
    fmt.Printf("Generated expression value: %v\n", val)
    return val
}

func (c *Compiler) handleFunctionCall(ctx parser.IFunctionCallContext) {
    fmt.Println("\nHandling function call")
    var fnName string
    if varCtx := ctx.Variable(); varCtx != nil {
        fnName = varCtx.GetText()
    }

    var args []value.Value
    for _, exprCtx := range ctx.AllExpression() {
        args = append(args, c.exprGen.GenerateExpression(exprCtx.(*parser.ExpressionContext)))
    }

    fmt.Printf("Function name: %s\n", fnName)
    fmt.Printf("Arguments count: %d\n", len(args))
    
    switch fnName {
    case "print":
        if len(args) == 0 {
            panic("print requires at least one argument")
        }
        fmt.Printf("Generating print call for value: %v\n", args[0])
        fmtName := fmt.Sprintf("fmt.str.%d", c.getUniqueID())
        formatStr := c.module.NewGlobalDef(fmtName, 
            constant.NewCharArrayFromString("%d\n\x00"))
        formatPtr := c.entry.NewBitCast(formatStr, types.I8Ptr)
        c.entry.NewCall(c.printf, formatPtr, args[0])
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
    fmt.Println("\n=== Finalizing IR ===")
    c.entry.NewRet(constant.NewInt(types.I32, 0))
    fmt.Println("Added final return instruction")
    return c.module.String(), nil
}