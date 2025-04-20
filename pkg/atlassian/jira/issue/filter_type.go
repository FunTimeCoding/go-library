package issue

import "slices"

func FilterType(
	v []*Issue,
	issueType ...string,
) []*Issue {
	var result []*Issue

	for _, i := range v {
		if slices.Contains(issueType, i.Type) {
			continue
		}

		result = append(result, i)
	}

	return result
}
