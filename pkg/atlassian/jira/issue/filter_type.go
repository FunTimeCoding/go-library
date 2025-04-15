package issue

import "slices"

func FilterType(
	v []*Issue,
	issueType ...string,
) []*Issue {
	var result []*Issue

	for _, element := range v {
		if slices.Contains(issueType, element.Type) {
			continue
		}

		result = append(result, element)
	}

	return result
}
