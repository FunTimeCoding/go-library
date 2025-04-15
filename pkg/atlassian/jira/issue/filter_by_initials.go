package issue

import "slices"

func ByInitials(
	v []*Issue,
	initials ...string,
) []*Issue {
	var result []*Issue

	for _, element := range v {
		if !slices.Contains(initials, element.Initials) {
			continue
		}

		result = append(result, element)
	}

	return result
}
