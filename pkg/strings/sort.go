package strings

import "sort"

func Sort(
	v []string,
	ascending bool,
) {
	sort.Strings(v)

	if !ascending {
		Reverse(v)
	}
}
