package variable_naming

import "go/types"

func isByteSlice(t types.Type) bool {
	s, okay := t.(*types.Slice)

	if !okay {
		return false
	}

	return isBasicKind(s.Elem(), types.Byte)
}
