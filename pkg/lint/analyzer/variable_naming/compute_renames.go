package variable_naming

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

// ComputeRenames returns the renames needed for a function's local variables.
func ComputeRenames(
	p *packages.Package,
	fn *ast.FuncDecl,
) []Rename {
	variables := collectVariables(p.TypesInfo, fn.Body)

	if len(variables) == 0 {
		return nil
	}

	assignments := assignLetters(variables)
	var result []Rename

	for _, v := range variables {
		if !isEligible(v) {
			continue
		}

		ideal, ok := assignments[v.ident]

		if !ok {
			continue
		}

		if v.ident.Name == ideal {
			continue
		}

		o := p.TypesInfo.ObjectOf(v.ident)

		if o == nil {
			continue
		}

		result = append(
			result,
			Rename{
				Object:  o,
				NewName: ideal,
			},
		)
	}

	return result
}
