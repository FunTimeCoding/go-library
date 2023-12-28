package contains

import "slices"

func All(
	expected []string,
	input []string,
) bool {
	for _, element := range expected {
		if !slices.Contains(input, element) {
			return false
		}
	}

	return true
}
