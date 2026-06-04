package palette

func characterBonus(
	previous rune,
	current rune,
) int {
	if isWhitespace(previous) {
		return bonusWhitespace
	}

	if isLower(previous) && isUpper(current) {
		return bonusCamelCase
	}

	if isDelimiter(previous) {
		return bonusWordBoundary
	}

	return 0
}
