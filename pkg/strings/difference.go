package strings

// difference returns the elements in a that are not in b
func difference(
	a []string,
	b []string,
) []string {
	buffer := make(map[string]struct{}, len(b))

	for _, x := range b {
		buffer[x] = struct{}{}
	}

	result := make([]string, 0)

	for _, x := range a {
		if _, okay := buffer[x]; !okay {
			result = append(result, x)
		}
	}

	return result
}
