package strings

func IndexOf(
	value string,
	s []string,
) int {
	for k, v := range s {
		if value == v {
			return k
		}
	}

	return -1
}
