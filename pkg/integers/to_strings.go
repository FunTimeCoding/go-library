package integers

import "strconv"

func ToStrings(v []int) []string {
	result := make([]string, len(v))

	for i, e := range v {
		result[i] = strconv.Itoa(e)
	}

	return result
}
