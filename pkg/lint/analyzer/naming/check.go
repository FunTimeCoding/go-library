package naming

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func check(
	pass *analysis.Pass,
	ident *ast.Ident,
	obj types.Object,
) {
	_, isVar := obj.(*types.Var)

	for _, segment := range segments(ident.Name) {
		if noSuggestion[segment] {
			pass.Reportf(
				ident.Pos(),
				"avoid %q in name, use a more specific term",
				segment,
			)

			return
		}

		alts, ok := suggestions[segment]

		if !ok {
			continue
		}

		var applicable []string

		for _, alt := range alts {
			if len(alt) > 1 || isVar {
				applicable = append(applicable, alt)
			}
		}

		if len(applicable) == 0 {
			pass.Reportf(
				ident.Pos(),
				"avoid %q in name, use a more specific term",
				segment,
			)

			return
		}

		for _, alt := range applicable {
			if containsSegment(ident.Name, alt) {
				return
			}
		}

		pass.Reportf(
			ident.Pos(),
			"%s",
			formatMessage(applicable, segment, ident.Name),
		)

		return
	}
}
