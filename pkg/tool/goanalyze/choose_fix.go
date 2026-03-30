package goanalyze

func chooseFix(
	name string,
	applicable []string,
) string {
	if len(segments(name)) > 1 {
		for _, a := range applicable {
			if len(a) > 1 {
				return a
			}
		}
	}

	return applicable[0]
}
