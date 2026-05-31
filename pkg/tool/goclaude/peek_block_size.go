package goclaude

func peekBlockSize(lines int) int {
	n := lines / 400

	if n < 3 {
		n = 3
	}

	if n > 10 {
		n = 10
	}

	return n
}
