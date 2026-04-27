package variable_naming

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func checkFunction(
	p *analysis.Pass,
	f *ast.FuncDecl,
) {
	variables := collectVariables(p.TypesInfo, f.Body)
	collectFromParams(p.TypesInfo, f, &variables)
	collectFromReceiver(p.TypesInfo, f, &variables)
	applyParameterExemptions(variables)

	if len(variables) == 0 {
		return
	}

	for i := range variables {
		variables[i].descendantNames = collectDescendantNames(
			p.TypesInfo,
			f,
			variables[i].ident,
		)
	}

	assignments := assignLetters(variables)

	for _, v := range variables {
		if !isEligible(v) {
			continue
		}

		ideal, okay := assignments[v.ident]

		if !okay {
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

	checkOkay(p, f)
}
