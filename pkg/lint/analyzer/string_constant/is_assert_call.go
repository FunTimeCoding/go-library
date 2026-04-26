package string_constant

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"strings"
)

func isAssertCall(
	p *analysis.Pass,
	call *ast.CallExpr,
) bool {
	s, ok := call.Fun.(*ast.SelectorExpr)

	if !ok {
		return false
	}

	i, ok := s.X.(*ast.Ident)

	if !ok {
		return false
	}

	o := p.TypesInfo.ObjectOf(i)

	if o == nil {
		return false
	}

	n, ok := o.(*types.PkgName)

	if !ok {
		return false
	}

	return strings.HasSuffix(n.Imported().Path(), assertPackageSuffix)
}
