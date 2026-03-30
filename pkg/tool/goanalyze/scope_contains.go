package goanalyze

import "go/types"

func scopeContains(scope *types.Scope, name string) bool {
	return scope.Lookup(name) != nil || childScopeContains(scope, name)
}
