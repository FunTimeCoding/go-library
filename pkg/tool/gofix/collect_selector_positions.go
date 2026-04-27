package gofix

import (
	"go/ast"
	"go/token"
)

func collectSelectorPositions(file *ast.File) map[token.Pos]bool {
	result := make(map[token.Pos]bool)
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			s, okay := n.(*ast.SelectorExpr)

			if okay {
				result[s.X.Pos()] = true
			}

			return true
		},
	)

	return result
}
