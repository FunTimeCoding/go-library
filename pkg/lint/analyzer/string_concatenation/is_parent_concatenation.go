package string_concatenation

import (
	"go/ast"
	"go/token"
)

func isParentConcatenation(parent ast.Node) bool {
	if parent == nil {
		return false
	}

	switch p := parent.(type) {
	case *ast.BinaryExpr:
		return p.Op == token.ADD
	case *ast.AssignStmt:
		return p.Tok == token.ADD_ASSIGN
	}

	return false
}
