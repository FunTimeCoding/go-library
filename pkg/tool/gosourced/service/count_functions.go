package service

import "go/ast"

func countFunctions(file *ast.File) int {
	var count int

	for _, d := range file.Decls {
		if _, okay := d.(*ast.FuncDecl); okay {
			count++
		}
	}

	return count
}
