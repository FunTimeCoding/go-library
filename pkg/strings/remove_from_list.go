package strings

func RemoveFromList(
	v []string,
	toRemove []string,
) []string {
	var result []string

	for i, value := range v {
		keep := true

		for _, remove := range toRemove {
			if value == remove {
				keep = false
			}
		}

		if keep {
			result = append(result, v[i])
		}
	}

	return result
}
