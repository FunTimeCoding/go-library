package expression

import "github.com/funtimecoding/go-library/pkg/strings/contains"

func (e *Expression) Check(s []string) bool {
	if contains.All(e.include, s) {
		if !contains.Any(e.exclude, s) {
			return true
		}
	}

	return false
}
