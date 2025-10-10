package slice

func StripEmpty(v []string) []string {
	var result []string

	for _, e := range v {
		if e != "" {
			result = append(result, e)
		}
	}

	return result
}
