package struct_literal

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func checkNew(
	p *analysis.Pass,
	call *ast.CallExpr,
	module string,
) {
	fun, okay := call.Fun.(*ast.Ident)

	if !okay || fun.Name != "new" {
		return
	}

	if len(call.Args) != 1 {
		return
	}

	sel, okay := call.Args[0].(*ast.SelectorExpr)

	if !okay {
		return
	}

	t := ownedType(p, sel, module)

	if t == nil {
		return
	}

	if suppress.IsSuppressed(p, call, "struct_literal") {
		return
	}

	p.Reportf(
		call.Pos(),
		"use a constructor function instead of new(%s.%s)",
		t.Pkg().Name(),
		t.Name(),
	)
}
