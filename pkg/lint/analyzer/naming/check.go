package naming

import (
	"fmt"
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

	_, isVariable := o.(*types.Var)

	for _, segment := range segments(ident.Name) {
		if noSuggestion[segment] {
			p.Report(analysis.Diagnostic{
				Pos:     ident.Pos(),
				End:     ident.End(),
				Message: fmt.Sprintf("avoid %q in name, use a more specific term", segment),
			})

			return
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
			p.Report(analysis.Diagnostic{
				Pos:     ident.Pos(),
				End:     ident.End(),
				Message: fmt.Sprintf("avoid %q in name, use a more specific term", segment),
			})

			return
		}

		for _, alternative := range applicable {
			if containsSegment(ident.Name, alternative) {
				return
			}
		}

		diagnostic := analysis.Diagnostic{
			Pos:     ident.Pos(),
			End:     ident.End(),
			Message: formatMessage(applicable, segment, ident.Name),
		}

		if o != nil && !ident.IsExported() {
			fix := chooseFix(ident.Name, applicable)
			replacement := replaceSegment(ident.Name, segment, fix)
			edits := buildEdits(p, o, segment, fix)

			if len(edits) > 0 {
				diagnostic.SuggestedFixes = []analysis.SuggestedFix{
					{
						Message:   fmt.Sprintf("rename to %s", replacement),
						TextEdits: edits,
					},
				}
			}
		}

		p.Report(diagnostic)

		return
	}
}
