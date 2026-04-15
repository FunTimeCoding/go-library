package variable_naming

import "go/types"

func isIntegerType(t types.Type) bool {
	b, ok := t.(*types.Basic)

	if !ok {
		return false
	}

	switch b.Kind() {
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:

		return true
	}

	return false
}
