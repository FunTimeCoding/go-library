package report

func spaces(number int) string {
	var result string

	for i := 0; i < number; i++ {
		result += "  "
	}

	return result
}
