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

	if scope == nil || scope.Lookup(replacement) == nil {
		return fix
	}

	if len(fix) > 1 {
		return ""
	}

	multiCharacter := firstMultiCharacter(applicable)

	if multiCharacter == "" {
		return ""
	}

	for _, r := range multiCharacter {
		letter := string(r)
		candidate := replaceSegment(name, segment, letter)

		if scope.Lookup(candidate) == nil {
			return letter
		}
	}

	return ""
}

func firstMultiCharacter(applicable []string) string {
	for _, a := range applicable {
		if len(a) > 1 {
			return a
		}
	}

	return ""
}
