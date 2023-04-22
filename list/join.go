package list

func Join[T any](slices [][]T) []T {
	var length int

	for _, s := range slices {
		length += len(s)
	}

	result := make([]T, length)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}
