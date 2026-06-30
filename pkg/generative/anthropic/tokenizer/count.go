package tokenizer

func (e *Encoder) Count(text string) int {
	count := 0
	position := 0

	for position < len(text) {
		match := e.trie.LongestMatch(text, position)

		if match == 0 {
			position++
		} else {
			position += match
		}

		count++
	}

	return count
}
