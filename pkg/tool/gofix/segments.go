package gofix

import (
	"strings"
	"unicode"
)

func segments(name string) []string {
	var result []string
	var current strings.Builder

	for _, part := range strings.Split(name, "_") {
		current.Reset()

		for i, r := range part {
			if i > 0 && unicode.IsUpper(r) {
				if s := current.String(); s != "" {
					result = append(result, s)
				}

				current.Reset()
			}

			current.WriteRune(unicode.ToLower(r))
		}

		if s := current.String(); s != "" {
			result = append(result, s)
		}
	}

	return result
}
