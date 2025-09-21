package page

import "sort"

func Names(v []*Page) []string {
	var result []string

	for _, g := range v {
		result = append(result, g.Name)
	}

	sort.Strings(result)

	return result
}
