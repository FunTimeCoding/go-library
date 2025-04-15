package strings

func IndexOfSkip(
	value string,
	s []string,
	skip int,
) int {
	for k, v := range s {
		if value == v {
			if skip > 0 {
				skip--

				continue
			}

			return k
		}
	}

	return -1
}
