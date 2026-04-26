package gofix

import (
	"go/ast"
	"go/types"
)

type violation struct {
	ident   *ast.Ident
	object  types.Object
	segment string
	fix     string
}
