package variable_naming

import "go/types"

func isErrorType(t types.Type) bool {
	if named, ok := t.(*types.Named); ok {
		return named.Obj().Pkg() == nil && named.Obj().Name() == "error"
	}

	return t.String() == "error"
}
