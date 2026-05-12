package parse

import (
	"go/ast"
	"go/token"
	"strings"
)

func StringValue(e ast.Expr) (string, bool) {
	l, okay := e.(*ast.BasicLit)

	if !okay || l.Kind != token.STRING {
		return "", false
	}

	return strings.Trim(l.Value, "\"`"), true
}
