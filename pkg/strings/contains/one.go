package contains

func One(
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
