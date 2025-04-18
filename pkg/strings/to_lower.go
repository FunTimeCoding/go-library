package strings

import "strings"

func ToLower(s []string) []string {
	var result []string

	for _, e := range s {
		result = append(result, strings.ToLower(e))
	}

	return result
}
