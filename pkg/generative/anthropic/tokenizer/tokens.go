package tokenizer

func (e *Encoder) Tokens(text string) []string {
	var result []string
	position := 0

	for position < len(text) {
		match := e.trie.LongestMatch(text, position)

		if match == 0 {
			result = append(result, text[position:position+1])
			position++
		} else {
			result = append(result, text[position:position+match])
			position += match
		}
	}

	return result
}
