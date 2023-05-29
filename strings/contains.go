package strings

func Contains(
	elements []string,
	search string,
) bool {
	for _, v := range elements {
		if search == v {
			return true
		}
	}

	return false
}
