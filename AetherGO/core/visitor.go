package core

import (
    "github.com/antlr4-go/antlr/v4"
    "FLUX/parser"
    //"strconv"
    //LLVMA "FLUX/AetherGO/LLVMA"
    "FLUX/AetherGO/utils"
)

type CustomVisitor struct {
	parser.BaseLua_grammar_antlr4Visitor
	Parser *parser.Lua_grammar_antlr4Parser
}

func (v *CustomVisitor) Visit(tree antlr.ParseTree) interface{} {
    return tree.Accept(v)
}

func (v *CustomVisitor) visitNode(node antlr.Tree) interface{} {
    switch ctx := node.(type) {
    case *antlr.TerminalNodeImpl:
        return ctx.GetText()
    case antlr.RuleContext:
        return v.VisitChildren(ctx)
    }
    return node
}

func (v *CustomVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
    if ctx == nil {
        utils.LogError("nil expression context in visitor")
        return nil
    }
    utils.LogInfo("=== NEW EXPRESSION PARSING ===")
    utils.LogInfo("Expression text: %s\n", ctx.GetText())
    utils.LogInfo("Child count: %d\n", ctx.GetChildCount())
    
    node := &ASTNode{Rule: "Expression"}
    
    defer func() {
        if len(node.Children) == 0 {
            utils.LogError("EMPTY EXPRESSION NODE CREATED - DEBUG INFO:")
            utils.LogError(" Original text: %s", ctx.GetText())
            utils.LogError(" Child count: %d", ctx.GetChildCount())
            for i := 0; i < ctx.GetChildCount(); i++ {
                child := ctx.GetChild(i)
                utils.LogError(" Child %d: %T (%s)", i, child, child.(antlr.ParseTree).GetText())
            }
        }
    }()

    for i := 0; i < ctx.GetChildCount(); {
        child := ctx.GetChild(i)
        utils.LogInfo("Processing child %d: %T (%s)\n", i, child, child.(antlr.ParseTree).GetText())
        
        if pexpr, ok := child.(*parser.PrimaryExpressionContext); ok {
            utils.LogInfo("- Found primary expression: %s\n", pexpr.GetText())
            currentExpr := v.Visit(pexpr)
            if node.Children == nil {
                node.Children = append(node.Children, currentExpr)
            } else {
                node.Children[len(node.Children)-1].(*ASTNode).Children = append(
                    node.Children[len(node.Children)-1].(*ASTNode).Children,
                    currentExpr,
                )
            }
            i++
            continue
        }
        if opGroup, ok := child.(*parser.OperatorGroupContext); ok {
            utils.LogInfo("- Found operator group: %s\n", opGroup.GetText())
            operatorNode := v.Visit(opGroup)
            binaryNode := &ASTNode{
                Rule:     "BinaryExpression",
                Children: []interface{}{},
            }
            if len(node.Children) > 0 {
                binaryNode.Children = append(binaryNode.Children, node.Children[len(node.Children)-1])
                node.Children = node.Children[:len(node.Children)-1]
            }
            
            binaryNode.Children = append(binaryNode.Children, operatorNode)
            node.Children = append(node.Children, binaryNode)
            i++
            if i < ctx.GetChildCount() {
                utils.LogInfo("- Processing right operand recursively\n")
                rightChild := ctx.GetChild(i).(antlr.ParseTree)
                rightOperand := v.Visit(rightChild)
                binaryNode.Children = append(binaryNode.Children, rightOperand)
                i++
            }
            continue
        }

        utils.LogWarning("UNHANDLED EXPRESSION CHILD TYPE: %T (%s)\n", child, child.(antlr.ParseTree).GetText())
        i++
    }
    
    utils.LogInfo("Final expression structure:\n")
    for _, child := range node.Children {
        utils.LogInfo("  %+v\n", child)
    }
    if len(node.Children) > 0 {
        switch node.Children[0].(type) {
        case *ASTNode:
            if node.Children[0].(*ASTNode).Rule == "Operator" {
                utils.LogError("MALFORMED EXPRESSION: Operator at start")
                utils.LogError(" Full context: %s", ctx.GetText())
            }
        }
    }

    return node
}

func (v *CustomVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
    node := &ASTNode{Rule: "PrimaryExpression"}
    if lit := ctx.Literal(); lit != nil {
        if num := lit.NUMBER(); num != nil {
            node.Children = append(node.Children, num.GetText())
        }
        if str := lit.STRING(); str != nil {
            node.Children = append(node.Children, str.GetText())
        }
    }
    return node
}

func (v *CustomVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
    root := &ASTNode{Rule: "Program"}
    for _, stmt := range ctx.AllStatement() {
        if node := v.Visit(stmt); node != nil {
            root.Children = append(root.Children, node)
        }
    }
    return root
}

func (v *CustomVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
    if assign := ctx.AssignStatement(); assign != nil {
        return v.Visit(assign)
    }
    return v.VisitChildren(ctx)
}

func (v *CustomVisitor) VisitChildren(node antlr.RuleNode) interface{} {
    astNode := &ASTNode{
        Rule: node.GetRuleContext().GetText(),
    }
    for _, child := range node.GetChildren() {
        astNode.Children = append(astNode.Children, v.visitNode(child))
    }
    return astNode
}

func (v *CustomVisitor) VisitLocalDeclaration(ctx *parser.LocalDeclarationContext) interface{} {
    node := &ASTNode{Rule: "LocalDeclaration"}
    node.Children = append(node.Children, 
        v.Visit(ctx.IDENTIFIER(0)),
        v.Visit(ctx.Expression()),
    )
    return node
}

func (v *CustomVisitor) VisitVariable(ctx *parser.VariableContext) interface{} {
    return &ASTNode{
        Rule:     "Variable",
        Children: []interface{}{ctx.GetText()},
    }
}

func (v *CustomVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
    node := &ASTNode{Rule: "FunctionCall"}
    if varCtx := ctx.Variable(); varCtx != nil {
        node.Children = append(node.Children, v.Visit(varCtx))
    }
    if args := ctx.GetChild(2); args != nil {
        if exprList, ok := args.(*parser.ExpressionListContext); ok {
            node.Children = append(node.Children, v.Visit(exprList))
        }
    }
    return node
}

func (v *CustomVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
    node := &ASTNode{Rule: "Arguments"}
    for _, expr := range ctx.AllExpression() {
        node.Children = append(node.Children, v.Visit(expr))
    }
    return node
}

func (v *CustomVisitor) VisitArithOp(ctx *parser.ArithOpContext) interface{} {
    return &ASTNode{
        Rule:     "Operator",
        Children: []interface{}{ctx.GetText()},
    }
}

func NewCustomVisitor(p *parser.Lua_grammar_antlr4Parser) *CustomVisitor {
	return &CustomVisitor{
		BaseLua_grammar_antlr4Visitor: parser.BaseLua_grammar_antlr4Visitor{},
		Parser: p,
	}
}