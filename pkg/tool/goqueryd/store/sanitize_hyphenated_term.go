package store

import "strings"

func sanitizeHyphenatedTerm(term string) string {
	parts := strings.Split(term, "-")
	var sanitized []string

	for _, p := range parts {
		t := sanitizeTerm(p)

		if t != "" {
			sanitized = append(sanitized, t)
		}
	}

	return strings.Join(sanitized, " ")
}
