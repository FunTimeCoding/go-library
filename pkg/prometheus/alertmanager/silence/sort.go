package silence

import "sort"

func Sort(v []*Silence) []*Silence {
	result := append([]*Silence(nil), v...)
	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].Start.After(*result[j].Start)
		},
	)

	return result
}
