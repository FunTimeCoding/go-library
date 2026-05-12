package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func File(path string) (*ast.File, *token.FileSet, error) {
	s := token.NewFileSet()
	f, e := parser.ParseFile(s, path, nil, parser.ParseComments)

	return f, s, e
}
