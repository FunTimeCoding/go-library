package variable_naming

import (
	"go/ast"
	"go/types"
)

func collectDescendantNames(
	y *types.Info,
	f *ast.FuncDecl,
	ident *ast.Ident,
) map[string]bool {
	result := map[string]bool{}
	block := enclosingBlock(f, ident)

	if block == nil {
		return result
	}

	end := ident.End()

	for _, s := range block.List {
		if s.End() <= end {
			continue
		}

		ast.Inspect(
			s,
			func(n ast.Node) bool {
				id, okay := n.(*ast.Ident)

				if okay && y.Defs[id] != nil {
					result[id.Name] = true
				}

				return true
			},
		)
	}

	return result
}
