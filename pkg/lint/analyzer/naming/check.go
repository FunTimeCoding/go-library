package naming

import (
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	skip := make(map[string]bool)

	for _, f := range p.Syntax {
		name := p.Fset.File(f.Pos()).Name()

		if isGenerated(name) || ast.IsGenerated(f) {
			skip[name] = true
		}
	}

	for _, file := range p.Syntax {
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				ident, okay := n.(*ast.Ident)

				if !okay {
					return true
				}

				if skip[p.Fset.File(ident.Pos()).Name()] {
					return true
				}

				o, isDefined := p.TypesInfo.Defs[ident]

				if !isDefined {
					return true
				}

				checkIdent(p, results, ident, o)

				return true
			},
		)
	}
}
