package main
import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"io/ioutil"
	"FLUX/parser" // PLEASE STOP ASKING ME IF THIS IS CORRECT OR NOT, FLUX IS THE PROJECT NAME SO IN GO, THIS IS CORRECT
	"strings"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"strconv"
	"path/filepath"
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
	parser := parser.NewLua_grammar_antlr4Parser(stream)
	visitor := NewCustomVisitor(parser)
	tree := parser.Program()
	result := visitor.Visit(tree)
	fmt.Println("AST Structure:")
	fmt.Println(formatAST(result, 0))

	// Generate and print LLVM IR
	module := generateIR(result)
	fmt.Println("\nGenerated LLVM IR:")
	fmt.Println(module)
}

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

// INIT AND DO NOTHING AFTER LOLOLOLOLOLOLOLOLOLOLOLLOOL ALR ILL IMPLEMENT IT LATER, cya
func generateIR(ast interface{}) *ir.Module {
	module := ir.NewModule()
	visited := make(map[string]bool)
	return generateIRRecursive(ast, module, visited, filepath.Dir(os.Args[1]))
}

func generateIRRecursive(ast interface{}, module *ir.Module, visited map[string]bool, basePath string) *ir.Module {
	printf := module.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true
	var mainFunc *ir.Func
	for _, f := range module.Funcs {
		if f.Name() == "main" {
			mainFunc = f
			break
		}
	}
	if mainFunc == nil {
		mainFunc = module.NewFunc("main", types.I32)
		entry := mainFunc.NewBlock("entry")
		entry.NewRet(constant.NewInt(types.I32, 0))
	}
	entry := mainFunc.Blocks[0]
	symbolTable := make(map[string]*ir.InstAlloca)
	createAlloca := func(name string) *ir.InstAlloca {
		if _, exists := symbolTable[name]; !exists {
			alloc := entry.NewAlloca(types.I32)
			symbolTable[name] = alloc
		}
		return symbolTable[name]
	}
	var genExpr func(interface{}) value.Value
	genExpr = func(node interface{}) value.Value {
		switch n := node.(type) {
		case map[string]interface{}:
			rule, ok := n["rule"].(string)
			if !ok {
				return nil
			}

			switch rule {
			case "number":
				if str, ok := n["nodes"].([]interface{})[0].(string); ok {
					return constant.NewInt(types.I32, int64(parseNumber(str)))
				}

			case "assignStatement":
				if asNodes, ok := n["nodes"].([]interface{}); ok && len(asNodes) > 2 {
					if varNode, ok := asNodes[0].(map[string]interface{}); ok {
						if varInnerNodes, ok := varNode["nodes"].([]interface{}); ok && len(varInnerNodes) > 0 {
							if varName, ok := varInnerNodes[0].(string); ok && varName != "" {
								alloc := createAlloca(varName)
								if value := genExpr(asNodes[2]); value != nil {
									entry.NewStore(value, alloc)
									return value
								}
							}
						}
					}
				}
				return nil

			case "functionDeclaration":
				if fnNameNode, ok := n["nodes"].([]interface{})[2].(map[string]interface{}); ok {
					if fnName, ok := fnNameNode["nodes"].([]interface{})[0].(string); ok {
						fn := module.NewFunc(fnName, types.I32,
							ir.NewParam("a", types.I32),
							ir.NewParam("b", types.I32),
						)
						block := fn.NewBlock("entry")
						sum := block.NewAdd(fn.Params[0], fn.Params[1])
						block.NewRet(sum)
					}
				}
				return nil

			case "functionCall":
				if fcNodes, ok := n["nodes"].([]interface{}); ok && len(fcNodes) > 0 {
					var fnName string
					if firstNode, ok := fcNodes[0].(map[string]interface{}); ok {
						if varNodes, ok := firstNode["nodes"].([]interface{}); ok && len(varNodes) > 0 {
							fnName, _ = varNodes[0].(string)
						}
					} else if directName, ok := fcNodes[0].(string); ok {
						fnName = directName
					}

					if fnName != "" {
						var args []value.Value
						if len(fcNodes) > 1 {
							if paramNode, ok := fcNodes[1].(map[string]interface{}); ok {
								if params, ok := paramNode["nodes"].([]interface{}); ok {
									for _, param := range params {
										args = append(args, genExpr(param))
									}
								}
							}
						}
						var fn *ir.Func
						for _, f := range module.Funcs {
							if f.Name() == fnName {
								fn = f
								break
							}
						}
						if fn == nil {
							params := make([]*ir.Param, len(args))
							for i := range params {
								params[i] = ir.NewParam("", types.I32)
							}
							fn = module.NewFunc(fnName, types.I32, params...)
						}
						return entry.NewCall(fn, args...)
					}
				}
				return nil

			case "operatorGroup":
				if op, ok := n["nodes"].([]interface{})[0].(string); ok {
					return &operatorWrapper{op: op}
				}
			case "expression":
				if len(n["nodes"].([]interface{})) > 1 {
					lhs := genExpr(n["nodes"].([]interface{})[0])
					opVal := genExpr(n["nodes"].([]interface{})[1])
					rhs := genExpr(n["nodes"].([]interface{})[2])
					if opWrapper, ok := opVal.(*operatorWrapper); ok {
						switch opWrapper.op {
						case "+":
							return entry.NewAdd(lhs.(value.Value), rhs.(value.Value))
						case "=":
							if lhsAlloca, isAlloca := lhs.(*ir.InstAlloca); isAlloca {
								entry.NewStore(rhs.(value.Value), lhsAlloca)
								return rhs.(value.Value)
							}
							if loadInst, isLoad := lhs.(*ir.InstLoad); isLoad {
								entry.NewStore(rhs.(value.Value), loadInst.Src)
								return rhs.(value.Value)
							}
						}
					}
				}
			case "variable":
				if nodes, ok := n["nodes"].([]interface{}); ok && len(nodes) > 0 {
					if varName, ok := nodes[0].(string); ok {
						if alloc, exists := symbolTable[varName]; exists {
							return entry.NewLoad(types.I32, alloc)
						}
					}
				}
			}
		case string:
			if num, err := strconv.Atoi(n); err == nil {
				return constant.NewInt(types.I32, int64(num))
			}
			if alloc, exists := symbolTable[n]; exists {
				return entry.NewLoad(types.I32, alloc)
			}
		}
		return nil
	}

	var processNode func(interface{})
	processNode = func(node interface{}) {
		switch n := node.(type) {
		case map[string]interface{}:
			if rule, ok := n["rule"].(string); ok {
				switch rule {
				case "functionDeclaration":
					if fnNameNode, ok := n["nodes"].([]interface{})[2].(map[string]interface{}); ok {
						if fnName, ok := fnNameNode["nodes"].([]interface{})[0].(string); ok {
							fn := module.NewFunc(fnName, types.I32,
								ir.NewParam("a", types.I32),
								ir.NewParam("b", types.I32),
							)
							block := fn.NewBlock("entry")
							sum := block.NewAdd(fn.Params[0], fn.Params[1])
							block.NewRet(sum)
						}
					}
				}
			}
			if nodes, ok := n["nodes"].([]interface{}); ok {
				for _, child := range nodes {
					switch child := child.(type) {
					case map[string]interface{}, string:
						processNode(child)
						genExpr(child)
					}
				}
			}

		case []interface{}:
			for _, child := range n {
				processNode(child)
			}
			
		case string:
			genExpr(n)
		}
	}
	processNode(ast)
	if len(entry.Insts) == 0 || entry.Insts[len(entry.Insts)-1].LLString() != "ret i32 0" {
		entry.NewRet(constant.NewInt(types.I32, 0))
	}
	
	return module
}

type operatorWrapper struct {
	op string
}

func (o *operatorWrapper) Type() types.Type {
	return types.Void
}

func (o *operatorWrapper) Ident() string {
	return o.op
}

func (o *operatorWrapper) String() string {
	return o.op
}