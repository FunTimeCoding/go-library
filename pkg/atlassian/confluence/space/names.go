package space

import "sort"

func Names(v []*Space) []string {
	var result []string

	for _, g := range v {
		result = append(result, g.Name)
	}

	sort.Strings(result)

	return result
}
