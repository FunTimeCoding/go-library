package naming

import (
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func buildEdits(
	p *analysis.Pass,
	o types.Object,
	segment string,
	fix string,
) []analysis.TextEdit {
	references := findReferences(p, o)
	var edits []analysis.TextEdit

	for _, reference := range references {
		edits = append(
			edits, analysis.TextEdit{
				Pos:     reference.Pos(),
				End:     reference.End(),
				NewText: []byte(replaceSegment(reference.Name, segment, fix)),
			},
		)
	}

	return edits
}
