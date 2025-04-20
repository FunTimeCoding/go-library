package issue

func FilterLabel(
	v []*Issue,
	labels ...string,
) []*Issue {
	var result []*Issue

	for _, i := range v {
		if i.HasLabel(labels...) {
			continue
		}

		result = append(result, i)
	}

	return result
}
