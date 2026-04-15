package variable_naming

import (
	"go/ast"
	"go/types"
)

func collectFromRange(
	info *types.Info,
	s *ast.RangeStmt,
	result *[]typedVariable,
) {
	for _, expr := range []ast.Expr{s.Key, s.Value} {
		if expr == nil {
			continue
		}

		ident, ok := expr.(*ast.Ident)

		if !ok || ident.Name == "_" {
			continue
		}

		o := info.ObjectOf(ident)

		if o == nil {
			continue
		}

		*result = append(
			*result,
			typedVariable{
				ident:      ident,
				typ:        o.Type(),
				precedence: typePrecedence(o.Type()),
			},
		)
	}
}
