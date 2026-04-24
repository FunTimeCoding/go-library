package file_identity

import (
	"go/ast"
	"go/token"
)

type identity struct {
	name   string
	position    token.Pos
	method *ast.FuncDecl
}
