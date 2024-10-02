package integers

import "strconv"

func ToStrings(v []int) []string {
	result := make([]string, len(v))

	for i, element := range v {
		result[i] = strconv.Itoa(element)
	}

	return result
}
