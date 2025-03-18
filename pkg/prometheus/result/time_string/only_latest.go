package time_string

func OnlyLatest(v []*Result) []*Result {
	var result []*Result
	sorted := make(map[string]*Result)

	for _, element := range v {
		value, exists := sorted[element.Metric]

		if exists {
			if value.Time.Before(element.Time) {
				sorted[element.Metric] = element
			}
		} else {
			sorted[element.Metric] = element
		}
	}

	for _, element := range sorted {
		result = append(result, element)
	}

	return result
}
