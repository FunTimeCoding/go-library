package strings

func DeleteDuplicates(values []string) []string {
	keys := make(map[string]bool)
	result := make([]string, 0)

	for _, value := range values {
		if _, okay := keys[value]; !okay {
			keys[value] = true
			result = append(result, value)
		}
	}

	return result
}
