package strings

func LinesAfter(
	input []string,
	match string,
) []string {
	var result []string
	var found bool

	for _, element := range input {
		if found {
			result = append(result, element)
		}

		if element == match {
			found = true
		}
	}

	return result
}
