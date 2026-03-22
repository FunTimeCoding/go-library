package string_concatenation

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func checkBinary(
	p *analysis.Pass,
	e *ast.BinaryExpr,
) {
	t := p.TypesInfo.TypeOf(e)

	if t == nil {
		return
	}

	b, ok := t.Underlying().(*types.Basic)

	if !ok || b.Kind() != types.String {
		return
	}

	if suppress.IsSuppressed(p, e, "string_concat") {
		return
	}

	p.Reportf(
		e.Pos(),
		"use fmt.Sprintf instead of string concatenation",
	)
}
