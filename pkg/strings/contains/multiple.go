package contains

import "github.com/funtimecoding/go-library/pkg/strings"

func Multiple(
	expected []string,
	input []string,
) bool {
	for _, element := range expected {
		if !strings.Contains(input, element) {
			return false
		}
	}

	return true
}
