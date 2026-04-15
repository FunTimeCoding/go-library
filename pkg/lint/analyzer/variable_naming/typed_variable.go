package variable_naming

import (
	"go/ast"
	"go/types"
)

type typedVariable struct {
	ident      *ast.Ident
	typ        types.Type
	precedence int
}
