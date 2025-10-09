package space

import "sort"

func Sort(v []*Space) []*Space {
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
