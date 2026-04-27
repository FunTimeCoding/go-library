package variable_naming

import (
	"go/ast"
	"go/types"
)

type typedVariable struct {
	ident           *ast.Ident
	typ             types.Type
	precedence      int
	scopedNames     map[string]bool
	descendantNames map[string]bool
	kind            variableKind
	exempt          bool
}
