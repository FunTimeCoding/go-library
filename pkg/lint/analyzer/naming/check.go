package naming

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func check(
	p *analysis.Pass,
	ident *ast.Ident,
	o types.Object,
) {
	if isInterfaceMethod(p, o) {
		return
	}

	_, isVar := o.(*types.Var)

	for _, segment := range segments(ident.Name) {
		if noSuggestion[segment] {
			p.Reportf(
				ident.Pos(),
				"avoid %q in name, use a more specific term",
				segment,
			)

			return
		}

		alts, hasSuggestion := suggestions[segment]

		if !hasSuggestion {
			continue
		}

		var applicable []string

		for _, alt := range alts {
			if len(alt) > 1 || isVar {
				applicable = append(applicable, alt)
			}
		}

		if len(applicable) == 0 {
			p.Reportf(
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

		p.Reportf(
			ident.Pos(),
			"%s",
			formatMessage(applicable, segment, ident.Name),
		)

		return
	}
}
