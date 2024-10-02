package strings

import "strings"

func ToLower(s []string) []string {
	var result []string

	for _, element := range s {
		result = append(result, strings.ToLower(element))
	}

	return result
}
