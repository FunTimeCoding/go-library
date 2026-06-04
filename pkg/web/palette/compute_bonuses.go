package palette

func computeBonuses(text []rune) []int {
	bonuses := make([]int, len(text))
	previous := ' '

	for i, current := range text {
		bonuses[i] = characterBonus(previous, current)
		previous = current
	}

	return bonuses
}
