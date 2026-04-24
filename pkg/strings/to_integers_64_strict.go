package strings

func ToIntegers64Strict(s []string) []int64 {
	result := make([]int64, len(s))

	for i, l := range s {
		result[i] = ToInteger64Strict(l)
	}

	return result
}
