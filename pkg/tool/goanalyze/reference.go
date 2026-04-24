package goanalyze

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type reference struct {
	ident *ast.Ident
	p     *packages.Package
}
