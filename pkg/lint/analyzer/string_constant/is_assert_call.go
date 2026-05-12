package string_constant

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

func isAssertCall(
	p *packages.Package,
	call *ast.CallExpr,
) bool {
	s, okay := call.Fun.(*ast.SelectorExpr)

	if !okay {
		return false
	}

	i, okay := s.X.(*ast.Ident)

	if !okay {
		return false
	}

	o := p.TypesInfo.ObjectOf(i)

	if o == nil {
		return false
	}

	n, okay := o.(*types.PkgName)

	if !okay {
		return false
	}

	return strings.HasSuffix(n.Imported().Path(), assertPackageSuffix)
}
