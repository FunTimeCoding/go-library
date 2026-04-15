package call_format

import (
	"go/ast"
	"go/token"
)

func lineLength(
	fileSet *token.FileSet,
	call *ast.CallExpr,
) int {
	end := fileSet.Position(call.End())

	return end.Column - 1
}
