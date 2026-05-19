package store

func parsePlainTerm(
	s string,
	i *int,
) string {
	start := *i

	for *i < len(s) && !isWhitespace(s[*i]) && s[*i] != '"' {
		*i++
	}

	return s[start:*i]
}
