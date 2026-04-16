package call_format

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func check(
	p *analysis.Pass,
	call *ast.CallExpr,
) {
	if len(call.Args) < 2 {
		return
	}

	if suppress.IsSuppressed(p, call, "call_format") {
		return
	}

	if !IsViolation(p.Fset, call) {
		return
	}

	openLine := p.Fset.Position(call.Lparen).Line
	closeLine := p.Fset.Position(call.Rparen).Line

	if openLine == closeLine {
		p.Reportf(
			call.Pos(),
			"call exceeds %d characters; split arguments to separate lines",
			maxLineLength,
		)
	} else {
		reportMultiLine(p, call)
	}
}
