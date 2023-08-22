package strings

func IndexOf(
	element string,
	elements []string,
) int {
	for key, value := range elements {
		if element == value {
			return key
		}
	}

	return -1
}
