package strings

func MustToIntegers(s []string) []int {
	result := make([]int, len(s))

	for i, l := range s {
		result[i] = MustToInteger(l)
	}

	return result
}
