package goclaude

func peekContextDepth(lines int) int {
	m := lines / 1500

	if m < 1 {
		m = 1
	}

	if m > 5 {
		m = 5
	}

	return m
}
