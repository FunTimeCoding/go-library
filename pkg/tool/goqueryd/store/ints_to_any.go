package store

func intsToAny(values []int) []any {
	result := make([]any, len(values))

	for i, v := range values {
		result[i] = v
	}

	return result
}
