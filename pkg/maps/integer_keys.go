package maps

import "sort"

func IntegerKeys[V any](v map[int]V) []int {
	var result []int

	for k := range v {
		result = append(result, k)
	}

	sort.Ints(result)

	return result
}
