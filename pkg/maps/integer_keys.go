package maps

func IntegerKeys[V any](v map[int]V) []int {
	var result []int

	for k := range v {
		result = append(result, k)
	}

	return result
}
