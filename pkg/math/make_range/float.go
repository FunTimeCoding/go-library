package make_range

func Float(
	minimum int,
	maximum int,
) []float64 {
	result := make([]float64, maximum-minimum+1)

	for index := range result {
		result[index] = float64(minimum + index)
	}

	return result
}
