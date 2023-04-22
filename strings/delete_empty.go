package strings

func DeleteEmpty(values []string) []string {
	var result []string

	for _, value := range values {
		if value != "" {
			result = append(result, value)
		}
	}

	return result
}
