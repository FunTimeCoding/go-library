package issue

func FilterBlocked(v []*Issue) []*Issue {
	var result []*Issue

	for _, i := range v {
		if !i.Blocked() {
			result = append(result, i)
		}
	}

	return result
}
