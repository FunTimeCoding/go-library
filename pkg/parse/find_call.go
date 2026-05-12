package parse

import "go/ast"

func FindCall(e ast.Expr, p string, function string) *ast.CallExpr {
	c, okay := e.(*ast.CallExpr)

	if !okay {
		return nil
	}

	if matchesCall(c, p, function) {
		return c
	}

	s, okay := c.Fun.(*ast.SelectorExpr)

	if !okay {
		return nil
	}

	return FindCall(s.X, p, function)
}
