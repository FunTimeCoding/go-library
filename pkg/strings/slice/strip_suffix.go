package slice

import "strings"

func StripSuffix(
	v []string,
	suffix string,
) []string {
	var result []string

	for _, e := range v {
		if !strings.HasSuffix(e, suffix) {
			result = append(result, e)
		}
	}

	return result
}
