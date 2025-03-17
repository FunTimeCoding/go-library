package alert

import "sort"

func SortByAge(
	v []*Alert,
	ascending bool,
) []*Alert {
	result := append([]*Alert{}, v...)
	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			if ascending {
				return result[i].Start.Before(*result[j].Start)
			}

			return result[i].Start.After(*result[j].Start)
		},
	)

	return result
}
