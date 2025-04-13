package event

import "sort"

func Sort(v []*Event) []*Event {
	result := append([]*Event{}, v...)
	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].Create.After(*result[j].Create)
		},
	)

	return result
}
