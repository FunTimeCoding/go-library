package forbidden_call

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func checkCall(
	p *analysis.Pass,
	call *ast.CallExpr,
) {
	sel, isSel := call.Fun.(*ast.SelectorExpr)

	if !isSel {
		return
	}

	o, isUse := p.TypesInfo.Uses[sel.Sel]

	if !isUse {
		return
	}

	fn, isFunction := o.(*types.Func)

	if !isFunction {
		return
	}

	a := fn.Pkg()

	if a == nil || a.Path() != "os/exec" {
		return
	}

	if !banned[fn.Name()] {
		return
	}

	if suppress.IsSuppressed(p, call, "forbidden_call") {
		return
	}

	p.Reportf(
		call.Pos(),
		"use pkg/system/run instead of exec.%s",
		fn.Name(),
	)
}
