package lint

import "sort"

func SortImports(imports []string) {
	sort.Slice(
		imports,
		func(
			i int,
			j int,
		) bool {
			pathI := ImportPath(imports[i])
			pathJ := ImportPath(imports[j])

			return pathI < pathJ
		},
	)
}
