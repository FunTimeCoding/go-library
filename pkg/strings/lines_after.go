package strings

func LinesAfter(
	v []string,
	match string,
) []string {
	var result []string
	var found bool

	for _, e := range v {
		if found {
			result = append(result, e)
		}

		if e == match {
			found = true
		}
	}

	return result
}
