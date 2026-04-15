package variable_naming

import "unicode"

func extractLastWord(name string) string {
	last := 0

	for i, r := range name {
		if unicode.IsUpper(r) && i > 0 {
			last = i
		}
	}

	return name[last:]
}
