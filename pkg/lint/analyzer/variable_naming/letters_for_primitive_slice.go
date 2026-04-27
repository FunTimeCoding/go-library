package variable_naming

import "go/types"

func lettersForPrimitiveSlice(t types.Type) []string {
	s, okay := t.Underlying().(*types.Slice)

	if !okay {
		return nil
	}

	return lettersForType(s.Elem())
}
