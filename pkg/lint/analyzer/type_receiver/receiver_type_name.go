package type_receiver

import "go/ast"

func receiverTypeName(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		if i, ok := t.X.(*ast.Ident); ok {
			return i.Name
		}
	case *ast.IndexExpr:
		if i, ok := t.X.(*ast.Ident); ok {
			return i.Name
		}
	case *ast.IndexListExpr:
		if i, ok := t.X.(*ast.Ident); ok {
			return i.Name
		}
	}

	return ""
}
