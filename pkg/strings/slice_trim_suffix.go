package strings

import "strings"

func SliceTrimSuffix(
	input []string,
	suffix string,
) []string {
	var result []string

	for _, element := range input {
		result = append(result, strings.TrimSuffix(element, suffix))
	}

	return result
}
