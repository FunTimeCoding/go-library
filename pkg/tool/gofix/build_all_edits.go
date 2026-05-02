package gofix

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/segment"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func buildAllEdits(
	fileSet *token.FileSet,
	all []*packages.Package,
	violations []violation,
	r *results,
) []edit {
	var result []edit

	for _, v := range violations {
		path := fileSet.File(v.ident.Pos()).Name()

		if v.fix == "" {
			r.addBlocked(
				path,
				fmt.Sprintf("cannot rename %s (collision)", v.ident.Name),
			)

			continue
		}

		replacement := segment.ReplaceSegment(v.ident.Name, v.segment, v.fix)
		references := findAllReferences(all, v.object)
		r.add(
			path,
			fmt.Sprintf(
				"renamed %s → %s (%d references)",
				v.ident.Name,
				replacement,
				len(references),
			),
		)

		for _, ref := range references {
			newName := segment.ReplaceSegment(ref.ident.Name, v.segment, v.fix)
			result = append(
				result,
				edit{
					position: ref.ident.Pos(),
					end:      ref.ident.End(),
					newText:  newName,
				},
			)
		}
	}

	return result
}
