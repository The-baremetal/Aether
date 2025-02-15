package core

import (
    "fmt"
)

type ASTNode struct {
    Rule     string
    Children []interface{}
}

func FormatAST(node interface{}, indent int) string {
    if n, ok := node.(*ASTNode); ok {
        str := fmt.Sprintf("%*s%s\n", indent, "", n.Rule)
        for _, child := range n.Children {
            str += FormatAST(child, indent+2)
        }
        return str
    }
    return fmt.Sprintf("%*s%v\n", indent, "", node)
}
