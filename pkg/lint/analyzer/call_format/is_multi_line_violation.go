package call_format

import (
	"go/ast"
	"go/token"
)

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
