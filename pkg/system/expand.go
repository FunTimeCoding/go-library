package system

import "strings"

func Expand(path string) string {
	if strings.HasPrefix(path, Tilde) {
		return strings.Replace(path, Tilde, Home(), 1)
	}

	return path
}
