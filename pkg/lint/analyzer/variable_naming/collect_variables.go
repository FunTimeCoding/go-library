package variable_naming

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func collectVariables(
	p *analysis.Pass,
	body *ast.BlockStmt,
) []typedVariable {
	var result []typedVariable
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.AssignStmt:
				collectFromAssign(p, node, &result)
			case *ast.RangeStmt:
				collectFromRange(p, node, &result)
			}

			return true
		},
	)

	return result
}
