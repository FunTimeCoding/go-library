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

	diagnostic := analysis.Diagnostic{
		Pos:     ident.Pos(),
		End:     ident.End(),
		Message: segment.FormatMessage(r.Applicable, r.Segment, ident.Name),
	}

	if o != nil && !ident.IsExported() {
		fix := segment.ResolveFix(ident.Name, r.Segment, r.Applicable, o)

		if fix != "" {
			replacement := segment.ReplaceSegment(ident.Name, r.Segment, fix)
			edits := buildEdits(p, o, r.Segment, fix)

			if len(edits) > 0 {
				diagnostic.SuggestedFixes = []analysis.SuggestedFix{
					{
						Message: fmt.Sprintf(
							"rename to %s",
							replacement,
						),
						TextEdits: edits,
					},
				}
			}
		}
	}

	p.Report(diagnostic)
}
