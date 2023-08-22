package strings

import "strings"

func ToLower(input []string) []string {
	var result []string

	for _, element := range input {
		result = append(result, strings.ToLower(element))
	}

	return result
}
