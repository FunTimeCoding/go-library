package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func Source(
	name string,
	content string,
) (*ast.File, *token.FileSet, error) {
	s := token.NewFileSet()
	f, e := parser.ParseFile(s, name, content, parser.ParseComments)

	return f, s, e
}
