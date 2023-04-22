package strings

// common returns the elements in a that are in b
func common(
	a []string,
	b []string,
) []string {
	buffer := make(map[string]struct{}, len(b))

	for _, x := range b {
		buffer[x] = struct{}{}
	}

	result := make([]string, 0)

	for _, x := range a {
		if _, found := buffer[x]; found {
			result = append(result, x)
		}
	}

	return result
}
