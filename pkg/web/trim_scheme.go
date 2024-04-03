package web

import "strings"

func TrimScheme(s string) string {
	if strings.HasPrefix(s, SecureSchemePrefix) {
		return strings.TrimPrefix(s, SecureSchemePrefix)
	}

	if strings.HasPrefix(s, InsecureSchemePrefix) {
		return strings.TrimPrefix(s, InsecureSchemePrefix)
	}

	return s
}
