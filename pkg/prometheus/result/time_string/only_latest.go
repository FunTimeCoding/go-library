package time_string

func OnlyLatest(v []*Result) []*Result {
	var result []*Result
	sorted := make(map[string]*Result)

	for _, e := range v {
		value, exists := sorted[e.Metric]

		if exists {
			if value.Time.Before(e.Time) {
				sorted[e.Metric] = e
			}
		} else {
			sorted[e.Metric] = e
		}
	}

	for _, e := range sorted {
		result = append(result, e)
	}

	return result
}
