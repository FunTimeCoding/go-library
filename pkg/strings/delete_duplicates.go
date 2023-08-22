package strings

func DeleteDuplicates(values []string) []string {
	keys := make(map[string]bool)
	result := make([]string, 0)

	for _, value := range values {
		if _, found := keys[value]; !found {
			keys[value] = true
			result = append(result, value)
		}
	}

	return result
}
