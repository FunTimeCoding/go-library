package call_format

import (
	"go/ast"
	"go/token"
	"strings"
)

func callIndent(
	fileSet *token.FileSet,
	call *ast.CallExpr,
	source []byte,
) string {
	p := fileSet.Position(call.Pos())
	lineStart := p.Offset - (p.Column - 1)
	count := 0

	for i := lineStart; i < len(source) && source[i] == '\t'; i++ {
		count++
	}

	return strings.Repeat("\t", count)
}
