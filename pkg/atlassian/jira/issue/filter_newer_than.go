package issue

func FilterNewerThan(
	v []*Issue,
	hours int,
) []*Issue {
	var result []*Issue

	for _, element := range v {
		if element.NewerThanHours(hours) {
			continue
		}

		result = append(result, element)
	}

	return result
}
