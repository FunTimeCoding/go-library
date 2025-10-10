package slice

import "strings"

func Trim(
	v []string,
	set string,
) []string {
	var result []string

	for _, e := range v {
		result = append(result, strings.Trim(e, set))
	}

	return result
}
