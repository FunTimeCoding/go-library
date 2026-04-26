package gofix

import (
	"fmt"
	"golang.org/x/tools/go/packages"
)

func buildAllEdits(
	all []*packages.Package,
	violations []violation,
) []edit {
	var result []edit

	for _, v := range violations {
		if v.fix == "" {
			fmt.Printf(
				"%s: cannot auto-fix (collision)\n",
				v.ident.Name,
			)

			continue
		}

		replacement := replaceSegment(v.ident.Name, v.segment, v.fix)
		references := findAllReferences(all, v.object)
		fmt.Printf(
			"Renamed: %s → %s (%d references)\n",
			v.ident.Name,
			replacement,
			len(references),
		)

		for _, r := range references {
			newName := replaceSegment(r.ident.Name, v.segment, v.fix)
			result = append(
				result,
				edit{
					position: r.ident.Pos(),
					end:      r.ident.End(),
					newText:  newName,
				},
			)
		}
	}

	return result
}
