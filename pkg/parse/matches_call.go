package parse

import "go/ast"

func matchesCall(c *ast.CallExpr, p string, function string) bool {
	s, okay := c.Fun.(*ast.SelectorExpr)

	if !okay {
		return false
	}

	i, okay := s.X.(*ast.Ident)

	if !okay {
		return false
	}

	return i.Name == p && s.Sel.Name == function
}
