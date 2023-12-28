package contains

import "slices"

func Any(
	valid []string,
	input []string,
) bool {
	for _, element := range valid {
		if slices.Contains(input, element) {
			return true
		}
	}

	return false
}
