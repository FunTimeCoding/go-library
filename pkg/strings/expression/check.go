package expression

import (
	contains2 "github.com/funtimecoding/go-library/pkg/strings/contains"
)

func (e *Expression) Check(input []string) bool {
	if contains2.Multiple(e.include, input) {
		if !contains2.Any(e.exclude, input) {
			return true
		}
	}

	return false
}
