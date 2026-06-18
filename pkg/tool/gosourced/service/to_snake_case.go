package service

import (
	"strings"
	"unicode"
)

func toSnakeCase(name string) string {
	var b strings.Builder

	for i, r := range name {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteRune('_')
			}

			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}
