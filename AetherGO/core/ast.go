package core

import (
    "fm"
    "FLUX/AetherGO/types"
)

func FormatAST(node interface{}, indent int) string {
    if n, ok := node.(*types.ASTNode); ok {
        typeStr := types.NodeTypeToString(n.Type)
        str := fmt.Sprintf("%*s%s\n", indent, "", typeStr)
        for _, child := range n.Children {
            str += FormatAST(child, indent+2)
        }
        return str
    }
    return fmt.Sprintf("%*s%v\n", indent, "", node)
}
