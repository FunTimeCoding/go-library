package strings

func ContainsMultiple(
	expected []string,
	input []string,
) bool {
	for _, element := range expected {
		if !Contains(input, element) {
			return false
		}
	}

	return true
}
