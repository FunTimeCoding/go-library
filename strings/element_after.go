package strings

func ElementAfter(
	elements []string,
	search string,
) string {
	var result string
	position := IndexOf(search, elements)

	if position != -1 {
		result = elements[position+1]
	}

	return result
}
