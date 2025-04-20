package issue

func FilterNewerThan(
	v []*Issue,
	hours int,
) []*Issue {
	var result []*Issue

	for _, i := range v {
		if i.NewerThanHours(hours) {
			continue
		}

		result = append(result, i)
	}

	return result
}
