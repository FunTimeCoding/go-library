package string_concatenation

import (
	"go/ast"
	"go/token"
)

func isNestedConcatenation(stack []ast.Node) bool {
	if len(stack) < 2 {
		return false
	}

	switch p := stack[len(stack)-2].(type) {
	case *ast.BinaryExpr:
		return p.Op == token.ADD
	case *ast.AssignStmt:
		return p.Tok == token.ADD_ASSIGN
	}

	return false
}
