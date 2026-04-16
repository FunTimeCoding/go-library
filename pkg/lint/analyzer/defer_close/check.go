package defer_close

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func check(
	p *analysis.Pass,
	d *ast.DeferStmt,
) {
	call, ok := d.Call.Fun.(*ast.SelectorExpr)

	if !ok || call.Sel.Name != "Close" {
		return
	}

	if len(d.Call.Args) != 0 {
		return
	}

	receiverType := p.TypesInfo.TypeOf(call.X)

	if receiverType == nil {
		return
	}

	if !implementsCloser(receiverType) {
		return
	}

	p.Reportf(
		d.Pos(),
		"use defer errors.PanicClose(%s) instead of defer %s.Close()",
		types.ExprString(call.X),
		types.ExprString(call.X),
	)
}
