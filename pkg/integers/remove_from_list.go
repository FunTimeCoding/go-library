package integers

func RemoveFromList(
	elements []int,
	toRemove []int,
) []int {
	var result []int

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
