package strings

func CountRune(
	text string,
	search rune,
) int {
	count := 0

	for _, singleCharacter := range text {
		if singleCharacter == search {
			count++
		}
	}

	return count
}
