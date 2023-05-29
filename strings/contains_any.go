package strings

func ContainsAny(
	valid []string,
	input []string,
) bool {
	for _, element := range valid {
		if Contains(input, element) {
			return true
		}
	}

	return false
}
