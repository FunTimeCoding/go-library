package service

import "go/ast"

func remainingFunctionName(file *ast.File) string {
	for _, d := range file.Decls {
		if f, okay := d.(*ast.FuncDecl); okay {
			return f.Name.Name
		}
	}

	return ""
}
