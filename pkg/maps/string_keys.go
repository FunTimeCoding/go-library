package maps

func StringKeys[V any](v map[string]V) []string {
	var result []string

	for k := range v {
		result = append(result, k)
	}

	return result
}
