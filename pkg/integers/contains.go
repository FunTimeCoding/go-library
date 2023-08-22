package integers

func Contains(
	elements []int,
	element int,
) bool {
	for _, v := range elements {
		if element == v {
			return true
		}
	}

	return false
}
