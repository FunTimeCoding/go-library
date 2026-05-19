package store

import (
	"regexp"
	"strings"
)

var nonAlphanumeric = regexp.MustCompile(`[^\p{L}\p{N}'_]+`)

func sanitizeTerm(term string) string {
	return nonAlphanumeric.ReplaceAllString(
		strings.ToLower(term),
		"",
	)
}
