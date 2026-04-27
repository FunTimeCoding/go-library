package variable_naming

import "go/types"

func isFloatType(t types.Type) bool {
	b, okay := t.(*types.Basic)

	if !okay {
		return false
	}

	return b.Kind() == types.Float32 || b.Kind() == types.Float64
}
