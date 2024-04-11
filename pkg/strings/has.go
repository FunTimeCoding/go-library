package strings

import "strings"

func Has(
	s string,
	prefix string,
	suffix string,
) bool {
	if !strings.HasPrefix(s, prefix) {
		return false
	}

	if !strings.HasSuffix(s, suffix) {
		return false
	}

	return true
}
