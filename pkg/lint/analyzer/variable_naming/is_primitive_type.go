package variable_naming

import "go/types"

func isPrimitiveType(t types.Type) bool {
	basic, okay := deref(t).Underlying().(*types.Basic)

	if !okay {
		return false
	}

	switch basic.Kind() {
	case types.Bool,
		types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64,
		types.Float32, types.Float64,
		types.String:

		return true
	}

	return false
}
