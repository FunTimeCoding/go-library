package contains

import "github.com/funtimecoding/go-library/pkg/strings"

func Any(
	valid []string,
	input []string,
) bool {
	for _, element := range valid {
		if strings.Contains(input, element) {
			return true
		}
	}

	return false
}
