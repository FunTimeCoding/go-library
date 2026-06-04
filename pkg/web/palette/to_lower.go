package palette

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}

	return r
}
