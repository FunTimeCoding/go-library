package token

func CountSlice(v []string) int {
	var result int

	for _, s := range v {
		result += Count(s)
	}

	return result
}
