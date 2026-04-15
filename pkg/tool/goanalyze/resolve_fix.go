package goanalyze

import "go/types"

func resolveFix(
	name string,
	segment string,
	applicable []string,
	o types.Object,
) string {
	scope := o.Parent()
	fix := chooseFix(name, applicable)
	replacement := replaceSegment(name, segment, fix)

	if scope == nil || !scopeContains(scope, replacement) {
		return fix
	}

	if len(fix) > 1 {
		return ""
	}

	for _, a := range applicable {
		if len(a) == 1 && a != fix {
			candidate := replaceSegment(name, segment, a)

			if !scopeContains(scope, candidate) {
				return a
			}
		}
	}

	firstWord := firstWord(applicable)

	if firstWord != "" {
		for _, r := range firstWord {
			letter := string(r)
			candidate := replaceSegment(name, segment, letter)

			if !scopeContains(scope, candidate) {
				return letter
			}
		}
	}

	for c := byte('a'); c <= byte('z'); c++ {
		letter := string(c)
		candidate := replaceSegment(name, segment, letter)

		if !scopeContains(scope, candidate) {
			return letter
		}
	}

	return ""
}
