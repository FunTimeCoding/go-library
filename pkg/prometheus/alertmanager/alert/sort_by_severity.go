package alert

import (
	"sort"
)

func SortBySeverity(
	v []*Alert,
	severities []string,
) []*Alert {
	m := make(map[string][]*Alert)

	for _, severity := range severities {
		for _, element := range v {
			if element.Severity == severity {
				m[severity] = append(m[severity], element)
			}
		}
	}

	for _, severity := range severities {
		sort.SliceStable(
			m[severity],
			func(
				i int,
				j int,
			) bool {
				return m[severity][i].Start.After(*m[severity][j].Start)
			},
		)
	}

	var result []*Alert

	for _, value := range m {
		result = append(result, value...)
	}

	return result
}
