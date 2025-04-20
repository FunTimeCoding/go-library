package issue

import "slices"

func FilterStatus(
	v []*Issue,
	status ...string,
) []*Issue {
	var result []*Issue

	for _, i := range v {
		if slices.Contains(status, i.Status) {
			continue
		}

		result = append(result, i)
	}

	return result
}
