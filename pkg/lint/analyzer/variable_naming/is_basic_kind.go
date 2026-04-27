package variable_naming

import "go/types"

func isBasicKind(
	t types.Type,
	kind types.BasicKind,
) bool {
	b, okay := t.(*types.Basic)

	if !okay {
		return false
	}

	return b.Kind() == kind
}
