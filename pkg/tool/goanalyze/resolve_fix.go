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

	multiCharacter := firstMultiCharacter(applicable)

	if multiCharacter == "" {
		return ""
	}

	for _, r := range multiCharacter {
		letter := string(r)
		candidate := replaceSegment(name, segment, letter)

		if !scopeContains(scope, candidate) {
			return letter
		}
	}

	return ""
}

func scopeContains(
	scope *types.Scope,
	name string,
) bool {
	return scope.Lookup(name) != nil || childScopeContains(scope, name)
}
