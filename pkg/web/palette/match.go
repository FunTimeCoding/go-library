package palette

func Match(
	pattern string,
	text string,
) (int, []int) {
	if pattern == "" {
		return 0, nil
	}

	patternRunes := toLowerRunes(pattern)
	textRunes := toLowerRunes(text)
	originalRunes := []rune(text)

	if len(patternRunes) > len(textRunes) {
		return -1, nil
	}

	if !containsAll(patternRunes, textRunes) {
		return -1, nil
	}

	bonuses := computeBonuses(originalRunes)
	score, positions := dynamicMatch(patternRunes, textRunes, bonuses)

	return score, positions
}
