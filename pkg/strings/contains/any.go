package contains

import "slices"

func Any(
	input []string,
	valid []string,
) bool {
	for _, element := range valid {
		if slices.Contains(input, element) {
			return true
		}
	}

	return false
}
