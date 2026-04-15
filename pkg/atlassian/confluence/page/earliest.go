package page

import "regexp"

func earliest(
	a []int,
	b []int,
	patternA *regexp.Regexp,
	patternB *regexp.Regexp,
) ([]int, *regexp.Regexp) {
	if a == nil {
		return b, patternB
	}

	if b == nil {
		return a, patternA
	}

	if a[0] <= b[0] {
		return a, patternA
	}

	return b, patternB
}
