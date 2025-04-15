package integers

func Contains(
	v []int,
	e int,
) bool {
	for _, l := range v {
		if e == l {
			return true
		}
	}

	return false
}
