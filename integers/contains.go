package integers

func Contains(
	element int,
	elements []int,
) bool {
	for _, value := range elements {
		if element == value {
			return true
		}
	}

	return false
}
