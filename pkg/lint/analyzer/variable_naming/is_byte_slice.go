package variable_naming

import "go/types"

func isByteSlice(t types.Type) bool {
	s, ok := t.(*types.Slice)

	if !ok {
		return false
	}

	return isBasicKind(s.Elem(), types.Byte)
}
