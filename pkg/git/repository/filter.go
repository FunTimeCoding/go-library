package repository

import "slices"

func Filter(
	v []*Repository,
	exclude []string,
	override bool,
) []*Repository {
	if override || len(exclude) == 0 {
		return v
	}

	var result []*Repository

	for _, r := range v {
		if slices.Contains(exclude, r.Path) {
			continue
		}

		result = append(result, r)
	}

	return result
}
