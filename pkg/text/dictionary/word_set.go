package dictionary

func wordSet(words []string) map[string]bool {
	s := make(map[string]bool, len(words))

	for _, w := range words {
		s[w] = true
	}

	return s
}
