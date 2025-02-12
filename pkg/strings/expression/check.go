package expression

import "github.com/funtimecoding/go-library/pkg/strings/contains"

func (e *Expression) Check(s []string) bool {
	if contains.All(s, e.include) {
		if !contains.Any(s, e.exclude) {
			return true
		}
	}

	return false
}
