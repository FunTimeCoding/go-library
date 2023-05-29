package expression

import "github.com/funtimecoding/go-library/strings"

func (e *Expression) Evaluate(input []string) bool {
	if strings.ContainsMultiple(e.include, input) {
		if !strings.ContainsAny(e.exclude, input) {
			return true
		}
	}

	return false
}
