package slice_sum

func Float(
	input []float64,
	index int,
) float64 {
	var result float64
	length := len(input)

	if index < 0 {
		index = 0
	} else if index > length {
		index = length
	}

	for _, e := range input[:index] {
		result += e
	}

	return result
}
