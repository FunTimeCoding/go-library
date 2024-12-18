package locator

import "strings"

func IsSubdomain(s string) bool {
	return strings.Count(s, ".") > 1
}
