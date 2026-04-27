package struct_literal

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func checkUnary(
	p *analysis.Pass,
	expr *ast.UnaryExpr,
	module string,
) {
	lit, okay := expr.X.(*ast.CompositeLit)

	if !okay {
		return
	}

	sel, okay := lit.Type.(*ast.SelectorExpr)

	if !okay {
		return
	}

	t := ownedType(p, sel, module)

	if t == nil {
		return
	}

	if suppress.IsSuppressed(p, expr, "struct_literal") {
		return
	}

	p.Reportf(
		expr.Pos(),
		"use a constructor function instead of &%s.%s{}",
		t.Pkg().Name(),
		t.Name(),
	)
}
