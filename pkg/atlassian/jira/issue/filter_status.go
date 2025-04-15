package issue

import "slices"

func FilterStatus(
	v []*Issue,
	status ...string,
) []*Issue {
	var result []*Issue

	for _, element := range v {
		if slices.Contains(status, element.Status) {
			continue
		}

		result = append(result, element)
	}

	return result
}
