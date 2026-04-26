package gofix

import "go/types"

func scopeContains(
	s *types.Scope,
	name string,
) bool {
	return s.Lookup(name) != nil ||
		childScopeContains(s, name) ||
		parentScopeContains(s, name)
}
