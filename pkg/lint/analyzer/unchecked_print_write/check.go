package unchecked_print_write

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func check(
	p *analysis.Pass,
	call *ast.CallExpr,
) {
	selector, okay := call.Fun.(*ast.SelectorExpr)

	if !okay {
		return
	}

	o, okay := p.TypesInfo.Uses[selector.Sel]

	if !okay {
		return
	}

	f, okay := o.(*types.Func)

	if !okay {
		return
	}

	a := f.Pkg()

	if a == nil || a.Path() != "fmt" || f.Name() != "Fprintf" {
		return
	}

	if suppress.IsSuppressed(p, call, "unchecked_print_write") {
		return
	}

	p.Reportf(
		call.Pos(),
		"use writer.Print, errors.Printf, or check the error from fmt.Fprintf",
	)
}
