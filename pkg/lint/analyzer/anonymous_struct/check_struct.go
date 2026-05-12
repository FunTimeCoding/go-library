package anonymous_struct

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkStruct(
	p *packages.Package,
	results *output.Results,
	s *ast.StructType,
) {
	if s.Fields == nil || len(s.Fields.List) == 0 {
		return
	}

	if isEmbeddingWrapper(s) {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, s.Pos(), "anonymous_struct") {
		return
	}

	results.AddBlocked(
		p.Fset.Position(s.Pos()).Filename,
		"anonymous struct; extract a named type",
	)
}
