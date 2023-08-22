package strings

func RemoveFromList(
	elements []string,
	toRemove []string,
) []string {
	var result []string

	for i, value := range elements {
		keep := true

		for _, remove := range toRemove {
			if value == remove {
				keep = false
			}
		}

		if keep {
			result = append(result, elements[i])
		}
	}

	return result
}
