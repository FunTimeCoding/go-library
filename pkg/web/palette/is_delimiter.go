package palette

func isDelimiter(r rune) bool {
	return r == '/' || r == ':' || r == ';' || r == '|' || r == '-' || r == '_'
}
