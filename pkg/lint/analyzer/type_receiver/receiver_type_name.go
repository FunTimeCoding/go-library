package type_receiver

import "go/ast"

func receiverTypeName(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		if i, okay := t.X.(*ast.Ident); okay {
			return i.Name
		}
	case *ast.IndexExpr:
		if i, okay := t.X.(*ast.Ident); okay {
			return i.Name
		}
	case *ast.IndexListExpr:
		if i, okay := t.X.(*ast.Ident); okay {
			return i.Name
		}
	}

	return ""
}
