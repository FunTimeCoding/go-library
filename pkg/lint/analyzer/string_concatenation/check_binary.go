package string_concatenation

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkBinary(
	p *packages.Package,
	results *output.Results,
	e *ast.BinaryExpr,
) {
	t := p.TypesInfo.TypeOf(e)

	if t == nil {
		return
	}

	b, okay := t.Underlying().(*types.Basic)

	if !okay || b.Kind() != types.String {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, e.Pos(), "string_concat") {
		return
	}

	results.AddBlocked(
		p.Fset.Position(e.Pos()).Filename,
		"use fmt.Sprintf instead of string concatenation",
	)
}
