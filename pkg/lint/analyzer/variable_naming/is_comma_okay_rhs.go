package variable_naming

import (
	"go/ast"
	"go/token"
	"go/types"
)

func isCommaOkayRHS(y *types.Info, expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.TypeAssertExpr:
		return true
	case *ast.IndexExpr:
		t := y.TypeOf(e.X)

		if t == nil {
			return false
		}

		_, okay := t.Underlying().(*types.Map)

		return okay
	case *ast.UnaryExpr:
		if e.Op != token.ARROW {
			return false
		}

		t := y.TypeOf(e.X)

		if t == nil {
			return false
		}

		_, okay := t.Underlying().(*types.Chan)

		return okay
	}

	return false
}
