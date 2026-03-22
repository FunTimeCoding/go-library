package string_concatenation

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func checkAssign(
	p *analysis.Pass,
	s *ast.AssignStmt,
) {
	t := p.TypesInfo.TypeOf(s.Lhs[0])

	if t == nil {
		return
	}

	b, ok := t.Underlying().(*types.Basic)

	if !ok || b.Kind() != types.String {
		return
	}

	if suppress.IsSuppressed(p, s, "string_concat") {
		return
	}

	p.Reportf(
		s.Pos(),
		"use fmt.Sprintf instead of string concatenation",
	)
}
