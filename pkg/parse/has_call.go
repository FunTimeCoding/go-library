package parse

import "go/ast"

func HasCall(
	f *ast.File,
	importPath string,
	function string,
) bool {
	name, found := ImportName(f, importPath)

	if !found {
		return false
	}

	var result bool
	ast.Inspect(
		f,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if okay && matchesCall(c, name, function) {
				result = true

				return false
			}

			return true
		},
	)

	return result
}
