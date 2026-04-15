package variable_naming

import (
	"go/ast"
	"go/types"
)

func collectVariables(
	info *types.Info,
	body *ast.BlockStmt,
) []typedVariable {
	var result []typedVariable
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.AssignStmt:
				collectFromAssign(info, node, &result)
			case *ast.RangeStmt:
				collectFromRange(info, node, &result)
			}

			return true
		},
	)

	return result
}
