package core

import (
    "github.com/antlr4-go/antlr/v4"
    "FLUX/parser"
    "FLUX/AetherGO/types"
    "strings"
)

type CustomVisitor struct {
    *parser.BaseLua_grammar_antlr4Visitor
    Parser *parser.Lua_grammar_antlr4Parser
}

func NewCustomVisitor(p *parser.Lua_grammar_antlr4Parser) *CustomVisitor {
    return &CustomVisitor{
        BaseLua_grammar_antlr4Visitor: &parser.BaseLua_grammar_antlr4Visitor{},
        Parser: p,
    }
}

func (v *CustomVisitor) Visit(tree antlr.ParseTree) interface{} {
    return v.visitNode(tree)
}

func (v *CustomVisitor) visitNode(node antlr.Tree) interface{} {
	switch ctx := node.(type) {
	case *antlr.TerminalNodeImpl:
		return ctx.GetText()
	case antlr.RuleContext:
		children := make([]interface{}, 0)
		ruleIndex := ctx.GetRuleIndex()
		ruleName := v.Parser.GetRuleNames()[ruleIndex]
		
		for i := 0; i < ctx.GetChildCount(); i++ {
			child := ctx.GetChild(i)
			if child != nil {
				children = append(children, v.visitNode(child))
			}
		}
		
		return map[string]interface{}{
			"rule":  ruleName,
			"nodes": children,
		}
	default:
		return nil
	}
}

func ruleNameToNodeType(ruleName string) types.NodeType {
    switch ruleName {
    case "Program":
        return types.NodeProgram
    case "IfStatement": 
        return types.NodeControlFlow
    case "FunctionDeclaration":
        return types.NodeDecl
    case "VariableDeclaration":
        return types.NodeLocalDecl
    }
    lower := strings.ToLower(ruleName)
    switch {
    case strings.Contains(lower, "expression"):
        return types.NodeExpr
    case strings.Contains(lower, "literal"):
        return types.NodeLiteral
    default:
        return types.NodeStmt
    }
}