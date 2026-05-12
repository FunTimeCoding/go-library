package naming

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"github.com/funtimecoding/go-library/pkg/lint/segment"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkIdent(
	p *packages.Package,
	results *output.Results,
	ident *ast.Ident,
	o types.Object,
) {
	if isInterfaceMethod(o) {
		return
	}

	v, isVariable := o.(*types.Var)
	isField := isVariable && v.IsField()
	r := segment.Check(ident.Name, isVariable, isField)

	if r == nil {
		return
	}

	if r.Banned {
		results.AddBlocked(
			p.Fset.Position(ident.Pos()).Filename,
			fmt.Sprintf(
				"avoid %q in name, use a more specific term",
				r.Segment,
			),
		)

		return
	}

	results.AddBlocked(
		p.Fset.Position(ident.Pos()).Filename,
		segment.FormatMessage(
			r.Applicable,
			r.Segment,
			ident.Name,
		),
	)
}
