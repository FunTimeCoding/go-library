package ranges

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

	for _, element := range input[:index] {
		result += element
	}

	return result
}
