package variable_naming

import (
	"go/ast"
	"go/token"
	"go/types"
)

func collectFromAssign(
	info *types.Info,
	s *ast.AssignStmt,
	result *[]typedVariable,
) {
	if s.Tok != token.DEFINE {
		return
	}

	for _, lhs := range s.Lhs {
		ident, ok := lhs.(*ast.Ident)

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
