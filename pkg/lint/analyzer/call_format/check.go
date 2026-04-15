package call_format

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
)

const maxLineLength = 80

func check(p *analysis.Pass, call *ast.CallExpr) {
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

// IsViolation reports whether a call expression violates call formatting rules.
func IsViolation(
	fileSet *token.FileSet,
	call *ast.CallExpr,
) bool {
	if len(call.Args) < 2 {
		return false
	}

	openLine := fileSet.Position(call.Lparen).Line
	closeLine := fileSet.Position(call.Rparen).Line

	if openLine == closeLine {
		return lineLength(fileSet, call) > maxLineLength
	}

	return isMultiLineViolation(fileSet, call)
}

func isMultiLineViolation(
	fileSet *token.FileSet,
	call *ast.CallExpr,
) bool {
	openParenLine := fileSet.Position(call.Lparen).Line
	firstArgLine := fileSet.Position(call.Args[0].Pos()).Line

	if openParenLine == firstArgLine {
		return true
	}

	for i := 1; i < len(call.Args); i++ {
		previousEnd := fileSet.Position(call.Args[i-1].End()).Line
		currentStart := fileSet.Position(call.Args[i].Pos()).Line

		if previousEnd == currentStart {
			return true
		}
	}

	return false
}

func lineLength(
	fileSet *token.FileSet,
	call *ast.CallExpr,
) int {
	end := fileSet.Position(call.End())

	return end.Column - 1
}
