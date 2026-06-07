package strings

func MustToIntegers64(s []string) []int64 {
	result := make([]int64, len(s))

	for i, l := range s {
		result[i] = MustToInteger64(l)
	}

	return result
}
