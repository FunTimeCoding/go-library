package alert

func GroupByInstance(v []*Alert) map[string][]*Alert {
	result := make(map[string][]*Alert)

	for _, a := range v {
		i := a.Instance()

		if i == "" {
			continue
		}

		result[i] = append(result[i], a)
	}

	return result
}
