package integers

func Contains(
	elements []int,
	element int,
) bool {
	for _, value := range elements {
		if element == value {
			return true
		}
	}

	return false
}
