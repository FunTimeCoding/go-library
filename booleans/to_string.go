package booleans

func ToString(b bool) string {
	var result string

	if b {
		result = "1"
	} else {
		result = "0"
	}

	return result
}
