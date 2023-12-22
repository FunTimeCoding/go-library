package contains

func All(
	expected []string,
	input []string,
) bool {
	for _, element := range expected {
		if !One(input, element) {
			return false
		}
	}

	return true
}
