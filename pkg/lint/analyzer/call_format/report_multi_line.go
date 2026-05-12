package call_format

import (
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func reportMultiLine(
	p *packages.Package,
	results *output.Results,
	call *ast.CallExpr,
) {
	openParenLine := p.Fset.Position(call.Lparen).Line
	firstArgLine := p.Fset.Position(call.Args[0].Pos()).Line

	if openParenLine == firstArgLine {
		results.AddBlocked(
			p.Fset.Position(call.Args[0].Pos()).Filename,
			"each argument should be on its own line",
		)

		return
	}

	for i := 1; i < len(call.Args); i++ {
		previousEnd := p.Fset.Position(call.Args[i-1].End()).Line
		currentStart := p.Fset.Position(call.Args[i].Pos()).Line

		if previousEnd == currentStart {
			results.AddBlocked(
				p.Fset.Position(call.Args[i].Pos()).Filename,
				"each argument should be on its own line",
			)

			return
		}
	}
}
