package variable_naming

import "go/types"

func deref(t types.Type) types.Type {
	if p, okay := t.(*types.Pointer); okay {
		return p.Elem()
	}

	return t
}
