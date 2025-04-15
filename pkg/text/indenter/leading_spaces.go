package indenter

func leadingSpaces(line string) int {
	for i, e := range line {
		if e != ' ' {
			return i
		}
	}

	return len(line)
}
