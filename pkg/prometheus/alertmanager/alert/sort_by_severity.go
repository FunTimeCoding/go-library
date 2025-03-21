package alert

import "sort"

func SortBySeverity(
	v []*Alert,
	severities []string,
) []*Alert {
	var result []*Alert

	for _, s := range severities {
		var alerts []*Alert

		for _, a := range v {
			if a.Severity == s {
				alerts = append(alerts, a)
			}
		}

		sort.SliceStable(
			alerts,
			func(
				i int,
				j int,
			) bool {
				return alerts[i].Start.After(*alerts[j].Start)
			},
		)
		result = append(result, alerts...)
	}

	return result
}
