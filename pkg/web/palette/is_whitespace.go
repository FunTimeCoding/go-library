package palette

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
