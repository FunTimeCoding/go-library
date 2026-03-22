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
	lit, ok := expr.X.(*ast.CompositeLit)

	if !ok {
		return
	}

	sel, ok := lit.Type.(*ast.SelectorExpr)

	if !ok {
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
