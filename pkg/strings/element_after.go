package strings

func ElementAfter(
	v []string,
	search string,
) string {
	var result string

	if p := IndexOf(search, v); p != -1 {
		result = v[p+1]
	}

	return result
}
