package variable_naming

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func checkOkay(p *analysis.Pass, u *ast.FuncDecl) {
	forms := collectOkayForms(p.TypesInfo, u)

	if len(forms) == 0 {
		return
	}

	assignments := assignOkayNames(forms)

	for _, f := range forms {
		expected := assignments[f.ident]

		if f.ident.Name == expected {
			continue
		}

		p.Reportf(
			f.ident.Pos(),
			"variable %s should be named %s",
			f.ident.Name,
			expected,
		)
	}
}
