package variable_naming

import "go/ast"

func enclosingBlock(
	f *ast.FuncDecl,
	ident *ast.Ident,
) *ast.BlockStmt {
	var ancestors []ast.Node
	var found *ast.BlockStmt
	ast.Inspect(
		f,
		func(n ast.Node) bool {
			if n == nil {
				ancestors = ancestors[:len(ancestors)-1]

				return false
			}

			if n == ident {
				for i := len(ancestors) - 1; i >= 0; i-- {
					b, okay := ancestors[i].(*ast.BlockStmt)

					if okay {
						found = b

						break
					}
				}

				return false
			}

			ancestors = append(ancestors, n)

			return true
		},
	)

	if found == nil {
		return f.Body
	}

	return found
}
