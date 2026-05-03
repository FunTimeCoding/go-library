package anonymous_struct

import (
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/suppress"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func check(
	p *analysis.Pass,
	s *ast.StructType,
) {
	if s.Fields == nil || len(s.Fields.List) == 0 {
		return
	}

	if isEmbeddingWrapper(s) {
		return
	}

	if suppress.IsSuppressed(p, s, "anonymous_struct") {
		return
	}

	p.Reportf(
		s.Pos(),
		"anonymous struct; extract a named type",
	)
}
