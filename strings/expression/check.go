package expression

import "github.com/funtimecoding/go-library/strings/contains"

func (e *Expression) Check(input []string) bool {
	if contains.Multiple(e.include, input) {
		if !contains.Any(e.exclude, input) {
			return true
		}
	}

	return false
}
