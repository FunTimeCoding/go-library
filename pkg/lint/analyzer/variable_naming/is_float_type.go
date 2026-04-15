package variable_naming

import "go/types"

func isFloatType(t types.Type) bool {
	b, ok := t.(*types.Basic)

	if !ok {
		return false
	}

	return b.Kind() == types.Float32 || b.Kind() == types.Float64
}
