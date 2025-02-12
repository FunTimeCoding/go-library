package contains

import "slices"

func All(
	input []string,
	valid []string,
) bool {
	for _, element := range valid {
		if !slices.Contains(input, element) {
			return false
		}
	}

	return true
}
