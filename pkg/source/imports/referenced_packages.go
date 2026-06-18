package imports

import "go/ast"

func referencedPackages(node ast.Node) map[string]bool {
	result := make(map[string]bool)
	ast.Inspect(
		node,
		func(n ast.Node) bool {
			selector, okay := n.(*ast.SelectorExpr)

			if !okay {
				return true
			}

			ident, okay := selector.X.(*ast.Ident)

			if !okay {
				return true
			}

			result[ident.Name] = true

			return true
		},
	)

	return result
}
