package strings

import "strings"

func ToMap(
	input []string,
	separator string,
) map[string]string {
	result := make(map[string]string)

	for _, s := range input {
		if p := strings.Split(s, separator); len(p) == 2 {
			result[p[0]] = p[1]
		}
	}

	return result
}
