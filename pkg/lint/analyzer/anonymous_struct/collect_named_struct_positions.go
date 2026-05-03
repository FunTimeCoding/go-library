package anonymous_struct

import (
	"go/ast"
	"go/token"
)

func collectNamedStructPositions(file *ast.File) map[token.Pos]bool {
	result := map[token.Pos]bool{}
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			t, okay := n.(*ast.TypeSpec)

			if !okay {
				return true
			}

			if st, okay := t.Type.(*ast.StructType); okay {
				result[st.Pos()] = true
			}

			return true
		},
	)

	return result
}
