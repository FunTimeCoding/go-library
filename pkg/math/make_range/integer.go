package make_range

func Integer(
	minimum int,
	maximum int,
) []int {
	result := make([]int, maximum-minimum+1)

	for index := range result {
		result[index] = minimum + index
	}

	return result
}
