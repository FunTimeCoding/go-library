package slice

import "strings"

func TrimSuffix(
	v []string,
	suffix string,
) []string {
	var result []string

	for _, e := range v {
		result = append(result, strings.TrimSuffix(e, suffix))
	}

	return result
}
