package chunk

import (
	"go/ast"
	"go/token"
	"sort"
)

func collectBoundaries(
	set *token.FileSet,
	file *ast.File,
) []int {
	seen := map[int]bool{}

	for _, declaration := range file.Decls {
		position := set.Position(declaration.Pos()).Offset

		if position > 0 {
			seen[position] = true
		}
	}

	result := make([]int, 0, len(seen))

	for position := range seen {
		result = append(result, position)
	}

	sort.Ints(result)

	return result
}
