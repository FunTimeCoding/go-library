package strings

func IndexOfSkip(
	element string,
	elements []string,
	skip int,
) int {
	for key, value := range elements {
		if element == value {
			if skip > 0 {
				skip--

				continue
			}

			return key
		}
	}

	return -1
}
