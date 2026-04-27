package naming

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/segment"
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

	v, isVariable := o.(*types.Var)
	isField := isVariable && v.IsField()
	r := segment.Check(ident.Name, isVariable, isField)

	if r == nil {
		return
	}

	if r.Banned {
		p.Report(
			analysis.Diagnostic{
				Pos: ident.Pos(),
				End: ident.End(),
				Message: fmt.Sprintf(
					"avoid %q in name, use a more specific term",
					r.Segment,
				),
			},
		)

		return
	}

	p.Report(
		analysis.Diagnostic{
			Pos:     ident.Pos(),
			End:     ident.End(),
			Message: segment.FormatMessage(r.Applicable, r.Segment, ident.Name),
		},
	)
}
