package contains

import "slices"

func All(
	v []string,
	valid []string,
) bool {
	for _, e := range valid {
		if !slices.Contains(v, e) {
			return false
		}
	}

	return true
}
