package contains

import "slices"

func Any(
	input []string,
	valid []string,
) bool {
	for _, e := range valid {
		if slices.Contains(input, e) {
			return true
		}
	}

	return false
}
