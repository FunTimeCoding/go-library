package strings

import "strings"

func ToBoolean(s string) bool {
	s = strings.TrimSpace(s)

	if s == "0" {
		return false
	}

	if s == "1" {
		return true
	}

	return false
}
