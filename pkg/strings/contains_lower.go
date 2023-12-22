package strings

import "strings"

func ContainsLower(
	s string,
	sub string,
) bool {
	return strings.Contains(strings.ToLower(s), sub)
}
