package variable_naming

import "go/types"

func lettersForPrimitiveSlice(t types.Type) []string {
	s, ok := t.Underlying().(*types.Slice)

	if !ok {
		return nil
	}

	return lettersForType(s.Elem())
}
