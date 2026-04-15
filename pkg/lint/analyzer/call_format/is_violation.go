package call_format

import (
	"go/ast"
	"go/token"
)

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
