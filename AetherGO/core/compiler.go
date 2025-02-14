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
    "FLUX/AetherGO/utils"
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
    utils.LogInfo("\n=== Setting up parser ===\n")
    is := antlr.NewInputStream(input)
    lexer := parser.NewLua_grammar_antlr4Lexer(is)
    stream := antlr.NewCommonTokenStream(lexer, 0)
    c.parser = parser.NewLua_grammar_antlr4Parser(stream)
    tree := c.parser.Program()
    c.ast = tree
    utils.LogDebug("Raw AST structure:\n%s", antlr.TreesStringTree(c.ast, c.parser.GetRuleNames(), c.parser))
    
    c.visitor = NewCustomVisitor(c.parser)
    utils.LogInfo("AST structure printed above")
    return nil
}

func (c *Compiler) GenerateIR() error {
    utils.LogInfo("\n=== Generating IR ===\n")
    utils.LogDebug("Initializing module...\n")
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
    
    utils.LogInfo("Created main function: %s\n", c.mainFn.Name())
    utils.LogInfo("Created entry block: %s\n", c.entry.Name())
    utils.LogInfo("Initialized expression generator")
    
    return nil
}

func (c *Compiler) ProcessAST() error {
    utils.LogInfo("\n=== Processing AST ===\n")
    utils.LogDebug("Full AST structure:\n%s\n", antlr.TreesStringTree(c.ast, nil, nil))
    
    ctx := c.ast.(*parser.ProgramContext)
    stmtList := ctx.AllStatement()
    utils.LogInfo("Found %d statements\n", len(stmtList))
    
    for i, stmt := range stmtList {
        utils.LogInfo("\nProcessing statement %d (%T)\n", i, stmt)
        if assign := stmt.AssignStatement(); assign != nil {
            c.handleAssignment(assign)
        } else if expr := stmt.Expression(); expr != nil {
            val := c.handleExpression(expr.(antlr.ParserRuleContext))
            if val != nil && val.Type() != types.Void {
                if fc := expr.(*parser.ExpressionContext).PrimaryExpression().FunctionCall(); fc != nil {
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
        exprCtx := expr.(parser.IExpressionContext)
        val := c.exprGen.GenerateExpression(exprCtx)
        c.entry.NewStore(val, alloc)
    }
    
    c.variables[ident] = &Variable{
        Name: ident,
        Type: types.I32,
        Value: alloc,
    }
}

func (c *Compiler) handleExpression(ctx antlr.ParserRuleContext) value.Value {
    utils.LogInfo("\n=== HANDLING EXPRESSION ===")
    defer utils.LogInfo("=== EXPRESSION HANDLED ===\n")
    
    utils.LogInfo("Full context type: %T", ctx)
    utils.LogInfo("Context text: %s", ctx.GetText())
    
    exprCtx := ctx.(*parser.ExpressionContext)
    utils.LogInfo("Operator groups found: %d", len(exprCtx.AllOperatorGroup()))
    
    if ops := exprCtx.AllOperatorGroup(); len(ops) > 0 {
        utils.LogInfo("Processing binary operation chain")
        left := c.exprGen.GenerateExpression(exprCtx.Expression(0))
        utils.LogInfo("Initial left operand: %v", left)
        
        for i, opGroup := range ops {
            op := opGroup.GetText()
            utils.LogInfo("Processing operator %d: %s", i+1, op)
            utils.LogInfo("Fetching right operand at index %d", i+1)
            
            right := c.exprGen.GenerateExpression(exprCtx.Expression(i+1))
            utils.LogInfo("Right operand value: %v", right)
            
            left = c.exprGen.GenerateBinaryOp(left, right, op)
            utils.LogInfo("New combined value: %v", left)
        }
        return left
    }
    val := c.exprGen.GenerateExpression(exprCtx)
    utils.LogInfo("Generated expression value: %v\n", val)
    return val
}

func (c *Compiler) handleFunctionCall(ctx parser.IFunctionCallContext) {
    utils.LogInfo("\nHandling function call\n")
    var fnName string
    if varCtx := ctx.Variable(); varCtx != nil {
        fnName = varCtx.GetText()
    }

    var args []value.Value
    for _, exprCtx := range ctx.AllExpression() {
        args = append(args, c.exprGen.GenerateExpression(exprCtx.(*parser.ExpressionContext)))
    }

    utils.LogInfo("Function name: %s\n", fnName)
    utils.LogInfo("Arguments count: %d\n", len(args))
    
    switch fnName {
    case "print":
        if len(args) == 0 {
            utils.LogError("print requires at least one argument\n")
        }
        utils.LogInfo("Generating print call for value: %v\n", args[0])
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
    utils.LogInfo("\n=== Finalizing IR ===\n")
    c.entry.NewRet(constant.NewInt(types.I32, 0))
    utils.LogInfo("Added final return instruction\n")
    utils.LogSuccess("Compilation successful!\n")
    return c.module.String(), nil
}