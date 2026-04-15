package variable_naming

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func checkFunction(p *analysis.Pass, fn *ast.FuncDecl) {
	variables := collectVariables(p.TypesInfo, fn.Body)

	if len(variables) == 0 {
		return
	}

	assignments := assignLetters(variables)

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

		p.Reportf(
			v.ident.Pos(),
			"variable %s of type %s should be named %s",
			v.ident.Name,
			v.typ.String(),
			ideal,
		)
	}
}
