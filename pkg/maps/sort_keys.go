package maps

import "sort"

func SortKeys(m map[string]string) []string {
	var result []string

	for k := range m {
		result = append(result, k)
	}

	sort.Strings(result)

	return result
}
