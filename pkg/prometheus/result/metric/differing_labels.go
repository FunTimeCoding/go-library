package metric

import "slices"

func DifferingLabels(m []*Metric) []string {
	if len(m) <= 1 {
		return []string{}
	}

	var result []string
	first := m[0]

	for _, element := range m {
		for k, v := range element.Map() {
			if first.Label(k) != v {
				if !slices.Contains(result, k) {
					result = append(result, k)
				}
			}
		}
	}

	return result
}
