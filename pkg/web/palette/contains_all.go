package palette

func containsAll(
	pattern []rune,
	text []rune,
) bool {
	ti := 0

	for _, p := range pattern {
		found := false

		for ti < len(text) {
			if text[ti] == p {
				found = true
				ti++

				break
			}

			ti++
		}

		if !found {
			return false
		}
	}

	return true
}
