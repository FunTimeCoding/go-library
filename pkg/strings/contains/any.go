package contains

func Any(
	valid []string,
	input []string,
) bool {
	for _, element := range valid {
		if One(input, element) {
			return true
		}
	}

	return false
}
