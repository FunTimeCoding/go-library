package palette

import "unicode/utf8"

func toLowerRunes(s string) []rune {
	runes := make([]rune, 0, utf8.RuneCountInString(s))

	for _, r := range s {
		runes = append(runes, toLower(r))
	}

	return runes
}
