package call_format

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func reportMultiLine(p *analysis.Pass, call *ast.CallExpr) {
	openParenLine := p.Fset.Position(call.Lparen).Line
	firstArgLine := p.Fset.Position(call.Args[0].Pos()).Line

	if openParenLine == firstArgLine {
		p.Reportf(
			call.Args[0].Pos(),
			"each argument should be on its own line",
		)

		return
	}

	for i := 1; i < len(call.Args); i++ {
		previousEnd := p.Fset.Position(call.Args[i-1].End()).Line
		currentStart := p.Fset.Position(call.Args[i].Pos()).Line

		if previousEnd == currentStart {
			p.Reportf(
				call.Args[i].Pos(),
				"each argument should be on its own line",
			)

			return
		}
	}
}
