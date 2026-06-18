package imports

import "go/ast"

func sameAlias(
	a *ast.ImportSpec,
	b *ast.ImportSpec,
) bool {
	if a.Name == nil && b.Name == nil {
		return true
	}

	if a.Name == nil || b.Name == nil {
		return false
	}

	return a.Name.Name == b.Name.Name
}
