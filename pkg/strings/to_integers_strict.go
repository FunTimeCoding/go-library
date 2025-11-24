package strings

func ToIntegersStrict(s []string) []int {
	result := make([]int, len(s))

	for i, l := range s {
		result[i] = ToIntegerStrict(l)
	}

	return result
}
