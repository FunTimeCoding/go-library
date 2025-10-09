package page

import "sort"

func Sort(v []*Page) []*Page {
	sort.SliceStable(
		v,
		func(
			i int,
			j int,
		) bool {
			return v[i].Name < v[j].Name
		},
	)

	return v
}
