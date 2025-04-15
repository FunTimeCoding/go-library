package issue

import "sort"

func SortByScore(v []*Issue) []*Issue {
	result := append([]*Issue{}, v...)
	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].score > result[j].score
		},
	)

	return result
}
