package expression

import "github.com/funtimecoding/go-library/pkg/strings/contains"

func (e *Expression) Check(input []string) bool {
	if contains.All(e.include, input) {
		if !contains.Any(e.exclude, input) {
			return true
		}
	}

	return false
}
