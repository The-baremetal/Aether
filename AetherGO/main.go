package main
import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"io/ioutil"
	"FLUX/parser" // PLEASE STOP ASKING ME IF THIS IS CORRECT OR NOT, FLUX IS THE PROJECT NAME SO IN GO, THIS IS CORRECT
	"strings"
	"strconv"
	"FLUX/AetherGO/LLVMA"
)
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")

		return
	}
	fileName := os.Args[1]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	input := string(content)
	is := antlr.NewInputStream(input)
	lexer := parser.NewLua_grammar_antlr4Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLua_grammar_antlr4Parser(stream)
	visitor := NewCustomVisitor(p)
	tree := p.Program()
	result := visitor.Visit(tree)
	fmt.Println("AST Structure:")
	fmt.Println(formatAST(result, 0))
	m, err := LLVMA.InitModule()
	if err != nil {
		fmt.Println("error initializing module:", err)
		return
	}
	main, err := LLVMA.InitMain(m) // for now, this isn't optional and is forced, the flag for -no-main can be added later.
	if err != nil {
		fmt.Println("error initializing func:", err)
		return
	}
	_ = main
	fmt.Println(m)



} // im too lazy to move the ast generation to a separate file, so here it is.

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

func formatAST(result interface{}, indent int) string {
	var sb strings.Builder
	indentStr := strings.Repeat("  ", indent)
	
	switch val := result.(type) {
	case map[string]interface{}:
		sb.WriteString(fmt.Sprintf("%s%s {\n", indentStr, val["rule"]))
		if nodes, ok := val["nodes"].([]interface{}); ok {
			for _, node := range nodes {
				sb.WriteString(formatAST(node, indent+1))
			}
		}
		sb.WriteString(fmt.Sprintf("%s}\n", indentStr))
	case string:
		sb.WriteString(fmt.Sprintf("%s'%s'\n", indentStr, val))
	case []interface{}:
		for _, item := range val {
			sb.WriteString(formatAST(item, indent))
		}
	}
	return sb.String()
}
func parseNumber(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}