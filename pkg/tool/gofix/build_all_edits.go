package gofix

import (
	"fmt"
	"go/token"
	"golang.org/x/tools/go/packages"
)

type edit struct {
	position token.Pos
	end      token.Pos
	newText  string
}

func buildAllEdits(
	all []*packages.Package,
	violations []violation,
) []edit {
	var result []edit

	for _, v := range violations {
		replacement := replaceSegment(v.ident.Name, v.segment, v.fix)
		references := findAllReferences(all, v.object)

		fmt.Printf(
			"%s → %s (%d references)\n",
			v.ident.Name,
			replacement,
			len(references),
		)

		for _, r := range references {
			newName := replaceSegment(r.ident.Name, v.segment, v.fix)
			result = append(result, edit{
				position: r.ident.Pos(),
				end:      r.ident.End(),
				newText:  newName,
			})
		}
	}

	return result
}
