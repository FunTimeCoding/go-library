package maps

import "sort"

func StringKeys[V any](v map[string]V) []string {
	var result []string

	for k := range v {
		result = append(result, k)
	}

	sort.Strings(result)

	return result
}
