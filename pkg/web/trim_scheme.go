package web

import "strings"

func TrimScheme(s string) string {
	if strings.HasPrefix(s, SecureSchemePrefix) {
		s = strings.TrimPrefix(s, SecureSchemePrefix)
	}

	return s
}
