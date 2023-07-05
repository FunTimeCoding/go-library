package integers

import "strconv"

func ToStrings(input []int) []string {
	result := make([]string, len(input))

	for i, element := range input {
		result[i] = strconv.Itoa(element)
	}

	return result
}
