package strings

func ElementAfterSkip(
	v []string,
	search string,
	skip int,
) string {
	var result string
	position := IndexOfSkip(search, v, skip)

	if position != -1 {
		result = v[position+1]
	}

	return result
}
