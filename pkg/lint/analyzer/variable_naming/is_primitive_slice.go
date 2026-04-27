package variable_naming

import "go/types"

func isPrimitiveSlice(t types.Type) bool {
	s, okay := deref(t).Underlying().(*types.Slice)

	if !okay {
		return false
	}

	return isPrimitiveType(s.Elem())
}
