package issue

import "slices"

func OnlyInitials(
	v []*Issue,
	initials ...string,
) []*Issue {
	var result []*Issue

	for _, i := range v {
		if !slices.Contains(initials, i.Initials) {
			continue
		}

		result = append(result, i)
	}

	return result
}
