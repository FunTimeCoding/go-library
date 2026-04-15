package variable_naming

import "go/types"

func isBasicKind(t types.Type, kind types.BasicKind) bool {
	b, ok := t.(*types.Basic)

	if !ok {
		return false
	}

	return b.Kind() == kind
}
