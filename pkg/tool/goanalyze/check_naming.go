package goanalyze

import (
	"go/ast"
	"go/types"
)

func checkNaming(
	ident *ast.Ident,
	o types.Object,
) *violation {
	_, isVariable := o.(*types.Var)

	for _, segment := range segments(ident.Name) {
		if noSuggestion[segment] {
			return nil
		}

		alternatives, hasSuggestion := suggestions[segment]

		if !hasSuggestion {
			continue
		}

		var applicable []string

		for _, alternative := range alternatives {
			if len(alternative) > 1 || isVariable {
				applicable = append(applicable, alternative)
			}
		}

		if len(applicable) == 0 {
			return nil
		}

		for _, alternative := range applicable {
			if containsSegment(ident.Name, alternative) {
				return nil
			}
		}

		fix := resolveFix(ident.Name, segment, applicable, o)

		return &violation{
			ident:   ident,
			object:  o,
			segment: segment,
			fix:     fix,
		}
	}

	return nil
}
