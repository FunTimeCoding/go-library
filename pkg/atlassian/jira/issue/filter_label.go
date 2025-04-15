package issue

func FilterLabel(
	v []*Issue,
	labels ...string,
) []*Issue {
	var result []*Issue

	for _, element := range v {
		if element.HasLabel(labels...) {
			continue
		}

		result = append(result, element)
	}

	return result
}
