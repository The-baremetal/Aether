package main
import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"io/ioutil"
	"FLUX/parser"
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
	lexer := parser.Newlua_grammar_antlr4Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.Newlua_grammar_antlr4Parser(stream)
	tree := parser.Program()
	fmt.Println("Parse Tree:")
	fmt.Println(tree.ToStringTree(nil, parser))
}