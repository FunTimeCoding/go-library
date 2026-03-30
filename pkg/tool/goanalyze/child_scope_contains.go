package goanalyze

import "go/types"

func childScopeContains(
	scope *types.Scope,
	name string,
) bool {
	for i := range scope.NumChildren() {
		child := scope.Child(i)

		if child.Lookup(name) != nil {
			return true
		}

		if childScopeContains(child, name) {
			return true
		}
	}

	return false
}
