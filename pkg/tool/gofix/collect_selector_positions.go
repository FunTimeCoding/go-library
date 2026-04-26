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
			s, ok := n.(*ast.SelectorExpr)

			if ok {
				result[s.X.Pos()] = true
			}

			return true
		},
	)

	return result
}
