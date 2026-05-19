package store

import "unicode"

func isWhitespace(b byte) bool {
	return unicode.IsSpace(rune(b))
}
