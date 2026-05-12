package variable_naming

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkOkay(
	p *packages.Package,
	results *output.Results,
	u *ast.FuncDecl,
) {
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

		results.AddBlocked(
			p.Fset.Position(f.ident.Pos()).Filename,
			fmt.Sprintf(
				"variable %s should be named %s",
				f.ident.Name,
				expected,
			),
		)
	}
}
