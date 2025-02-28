package core

import (
	"fmt"
	"strings"
	"FLUX/AetherGO/types"
)

func FormatAST(node *types.ASTNode, indentLevel int) string {
	var sb strings.Builder
	indent := strings.Repeat("  ", indentLevel)
	if node == nil {
		return fmt.Sprintf("%s<nil>\n", indent)
	}
	sb.WriteString(fmt.Sprintf("%s%s", indent, types.NodeTypeToString(node.Type)))
	if node.Value != nil {
		sb.WriteString(fmt.Sprintf(" (%#v)", node.Value))
	}
	if len(node.Children) > 0 {
		sb.WriteString(" {\n")
		for _, child := range node.Children {
			sb.WriteString(FormatAST(child, indentLevel+1))
		}
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	} else {
		sb.WriteString("\n")
	}
	return sb.String()
}

