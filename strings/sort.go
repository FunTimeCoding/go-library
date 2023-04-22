package strings

import "sort"

func Sort(
	elements []string,
	ascending bool,
) {
	sort.Strings(elements)

	if !ascending {
		Reverse(elements)
	}
}
