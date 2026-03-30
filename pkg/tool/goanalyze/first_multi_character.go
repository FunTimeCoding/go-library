package goanalyze

func firstMultiCharacter(applicable []string) string {
	for _, a := range applicable {
		if len(a) > 1 {
			return a
		}
	}

	return ""
}
