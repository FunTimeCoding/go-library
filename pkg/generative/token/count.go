package token

func Count(s string) int {
	_, tokens := Encode(s)

	return len(tokens)
}
