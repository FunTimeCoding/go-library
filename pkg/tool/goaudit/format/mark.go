package format

func mark(v bool) string {
	if v {
		return "+"
	}

	return "-"
}
