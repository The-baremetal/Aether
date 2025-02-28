package core

import (
    "github.com/antlr4-go/antlr/v4"
    "FLUX/parser"
    "FLUX/AetherGO/types"
    "strings"
    "fmt"
)

type CustomVisitor struct {
    parser.BaseLua_grammar_antlr4Visitor
    Parser *parser.Lua_grammar_antlr4Parser
}

func NewCustomVisitor(p *parser.Lua_grammar_antlr4Parser) *CustomVisitor {
    return &CustomVisitor{
        BaseLua_grammar_antlr4Visitor: parser.BaseLua_grammar_antlr4Visitor{},
        Parser: p,
    }
}

func (v *CustomVisitor) Visit(tree antlr.ParseTree) interface{} {
    fmt.Println("Visit is called")
    fmt.Printf("Processing parse tree: %T\n", tree)
    result, err := v.visitNode(tree)
    if err != nil {
        fmt.Printf("Visitor error: %v\n", err)
        return &types.ASTNode{Type: types.NodeStmt}
    }
    return result
}

func (v *CustomVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
    node, _ := v.visitNode(ctx)
    return node
}

func (v *CustomVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
    node, _ := v.visitNode(ctx)
    return node
}

func (v *CustomVisitor) VisitLocalDecl(ctx *parser.LocalDeclContext) interface{} {
    node, _ := v.visitNode(ctx)
    return node
}

func (v *CustomVisitor) visitNode(node antlr.Tree) (*types.ASTNode, error) {
    switch ctx := node.(type) {
    case *antlr.TerminalNodeImpl:
        fmt.Printf("Processing terminal: %s\n", ctx.GetText())
        return &types.ASTNode{
            Type:  types.NodeLiteral,
            Value: ctx.GetText(),
        }, nil

    case *parser.ProgramContext:
        return v.handleProgram(ctx)
    case *parser.LocalDeclContext:
        return v.handleLocalDeclaration(ctx)

    case antlr.RuleContext:
        ruleName := v.Parser.GetRuleNames()[ctx.GetRuleIndex()]
        nodeType := ruleNameToNodeType(ruleName)
        
        astNode := &types.ASTNode{
            Type:     nodeType,
            Children: make([]*types.ASTNode, 0),
        }

        ruleValue, err := parseRuleValue(ruleName, ctx)
        if err != nil {
            if parserCtx, ok := ctx.(*antlr.BaseParserRuleContext); ok {
                return nil, fmt.Errorf("%s at line %d", err, parserCtx.GetStart().GetLine())
            }
            return nil, err
        }
        astNode.Value = ruleValue

        for i := 0; i < ctx.GetChildCount(); i++ {
            childNode, err := v.visitNode(ctx.GetChild(i))
            if err != nil {
                return nil, err
            }
            astNode.Children = append(astNode.Children, childNode)
        }
        return astNode, nil

    default:
        return nil, fmt.Errorf("unknown node type %T", node)
    }
}

func (v *CustomVisitor) handleProgram(ctx *parser.ProgramContext) (*types.ASTNode, error) {
    programNode := &types.ASTNode{
        Type:     types.NodeProgram,
        Value:    "program_root",
        Children: make([]*types.ASTNode, 0),
    }

    for _, stmt := range ctx.AllStatement() {
        stmtNode, err := v.visitNode(stmt)
        if err != nil {
            return nil, err
        }
        programNode.Children = append(programNode.Children, stmtNode)
    }

    return programNode, nil
}

func (v *CustomVisitor) handleLocalDeclaration(ctx *parser.LocalDeclContext) (*types.ASTNode, error) {
    var declType string
    var identifiers []string
    var values []*types.ASTNode
    switch {
    case ctx.TypeAnnotation() != nil:
        declType = "typed"
        for i := 0; i < len(ctx.AllTypeAnnotation()); i++ {
            typeCtx := ctx.TypeAnnotation(i)
            ident := ctx.IDENTIFIER(i).GetText()
            identifiers = append(identifiers, ident)
            typeNode, err := v.visitNode(typeCtx)
            if err != nil {
                return nil, err
            }
            identifiers = append(identifiers, typeNode.Value.(string))
        }
        
    case ctx.IDENTIFIER(0) != nil:
        declType = "untyped"
        for _, ident := range ctx.AllIDENTIFIER() {
            identifiers = append(identifiers, ident.GetText())
        }
        
    case ctx.FUNCTION() != nil:
        return v.handleLocalFunction(ctx)
        
    default:
        return nil, fmt.Errorf("invalid local declaration structure at line %d", ctx.GetStart().GetLine())
    }
    if exprList := ctx.ExpressionList(); exprList != nil {
        for _, expr := range exprList.AllExpression() {
            exprNode, err := v.visitNode(expr)
            if err != nil {
                return nil, err
            }
            values = append(values, exprNode)
        }
    }
    if len(values) > 0 && len(identifiers) != len(values) {
        return nil, fmt.Errorf("declaration count mismatch: %d identifiers vs %d values at line %d",
            len(identifiers), len(values), ctx.GetStart().GetLine())
    }
    node := &types.ASTNode{
        Type:  types.NodeLocalDecl,
        Value: declType,
        Children: []*types.ASTNode{
            {Type: types.NodeIdentList, Value: identifiers},
        },
    }
    if len(values) > 0 {
        node.Children = append(node.Children, &types.ASTNode{
            Type:     types.NodeExprList,
            Children: values,
        })
    }
    return node, nil
}

func parseRuleValue(ruleName string, ctx antlr.RuleContext) (interface{}, error) {
    parserCtx, ok := ctx.(*antlr.BaseParserRuleContext)
    if !ok {
        return nil, fmt.Errorf("invalid context type %T", ctx)
    }

    switch ruleName {
    case "FunctionDeclaration":
        if len(ctx.GetChildren()) == 0 {
            return nil, fmt.Errorf("FunctionDeclaration missing identifier at line %d", parserCtx.GetStart().GetLine())
        }
        if ident, ok := ctx.GetChild(0).(*antlr.TerminalNodeImpl); ok {
            return ident.GetText(), nil
        }
        return nil, fmt.Errorf("expected identifier for FunctionDeclaration at line %d", parserCtx.GetStart().GetLine())
    
    case "VariableDeclaration":
        if len(ctx.GetChildren()) == 0 {
            return nil, fmt.Errorf("VariableDeclaration missing identifier at line %d", parserCtx.GetStart().GetLine())
        }
        if ident, ok := ctx.GetChild(0).(*antlr.TerminalNodeImpl); ok {
            return ident.GetText(), nil
        }
        return nil, fmt.Errorf("expected identifier for VariableDeclaration at line %d", parserCtx.GetStart().GetLine())
    
    case "program":
        return "program_root", nil
    
    case "LocalDecl":
        if len(ctx.GetChildren()) < 2 {
            return nil, fmt.Errorf("LocalDecl missing identifier at line %d", parserCtx.GetStart().GetLine())
        }
        if ident, ok := ctx.GetChild(1).(*antlr.TerminalNodeImpl); ok {
            return ident.GetText(), nil
        }
        return nil, fmt.Errorf("expected identifier for LocalDecl at line %d", parserCtx.GetStart().GetLine())
    
    default:
        return nil, fmt.Errorf("parseRuleValue not implemented for %s", ruleName)
    }
}

func ruleNameToNodeType(ruleName string) types.NodeType {
    switch ruleName {
    case "program":
        return types.NodeProgram
    case "IfStatement", "WhileStatement", "RepeatStatement":
        return types.NodeControlFlow
    case "FunctionDeclaration":
        return types.NodeDecl
    case "VariableDeclaration", "LocalDecl":
        return types.NodeLocalDecl
    case "FunctionCall":
        return types.NodeFuncCall
    case "TableConstructor":
        return types.NodeExpr
    case "ReturnStatement":
        return types.NodeStmt
    case "AssignStatement":
        return types.NodeStmt
    case "ParameterList":
        return types.NodeArguments
    case "Expression":
        return types.NodeExpr
    case "NumberLiteral", "StringLiteral":
        return types.NodeLiteral
    }
    
    lower := strings.ToLower(ruleName)
    switch {
    case strings.Contains(lower, "expression"):
        return types.NodeExpr
    case strings.Contains(lower, "literal"):
        return types.NodeLiteral
    case strings.Contains(lower, "operator"):
        return types.NodeOperator
    default:
        return types.NodeStmt
    }
}