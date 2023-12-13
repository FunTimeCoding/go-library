package strings

func ElementAfterSkip(
	elements []string,
	search string,
	skip int,
) string {
	var result string
	position := IndexOfSkip(search, elements, skip)

	if position != -1 {
		result = elements[position+1]
	}

	return result
}
