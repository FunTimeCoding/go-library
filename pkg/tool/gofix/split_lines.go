package gofix

func splitLines(b []byte) []string {
	var result []string
	start := 0

	for i, c := range b {
		if c == '\n' {
			result = append(result, string(b[start:i]))
			start = i + 1
		}
	}

	if start < len(b) {
		result = append(result, string(b[start:]))
	}

	return result
}
