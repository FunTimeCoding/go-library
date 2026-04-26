package gofix

import "go/types"

func parentScopeContains(
	scope *types.Scope,
	name string,
) bool {
	for p := scope.Parent(); p != nil; p = p.Parent() {
		if p.Parent() == nil {
			return false
		}

		if p.Lookup(name) != nil {
			return true
		}
	}

	return false
}
