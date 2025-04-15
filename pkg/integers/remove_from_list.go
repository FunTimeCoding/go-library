package integers

func RemoveFromList(
	v []int,
	toRemove []int,
) []int {
	var result []int

	for i, a := range v {
		keep := true

		for _, remove := range toRemove {
			if a == remove {
				keep = false
			}
		}

		if keep {
			result = append(result, v[i])
		}
	}

	return result
}
