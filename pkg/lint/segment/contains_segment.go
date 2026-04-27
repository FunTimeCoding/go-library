package segment

func ContainsSegment(
	name string,
	target string,
) bool {
	for _, s := range Segments(name) {
		if len(target) == 1 {
			if s == target {
				return true
			}
		} else {
			if len(s) >= len(target) && s[:len(target)] == target {
				return true
			}
		}
	}

	return false
}
