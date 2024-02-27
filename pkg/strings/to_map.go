package strings

import "strings"

func ToMap(
	input []string,
	separator string,
) map[string]string {
	result := make(map[string]string)

	for _, s := range input {
		parts := strings.Split(s, separator)
		result[parts[0]] = parts[1]
	}

	return result
}
