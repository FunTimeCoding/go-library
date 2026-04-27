package naming

import (
	"github.com/funtimecoding/go-library/pkg/lint/segment"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func buildEdits(
	p *analysis.Pass,
	o types.Object,
	s string,
	fix string,
) []analysis.TextEdit {
	references := findReferences(p, o)
	var edits []analysis.TextEdit

	for _, reference := range references {
		edits = append(
			edits,
			analysis.TextEdit{
				Pos:     reference.Pos(),
				End:     reference.End(),
				NewText: []byte(segment.ReplaceSegment(reference.Name, s, fix)),
			},
		)
	}

	return edits
}
