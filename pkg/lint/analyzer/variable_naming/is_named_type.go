package variable_naming

import "go/types"

func isNamedType(
	t types.Type,
	packagePath string,
	name string,
) bool {
	t = deref(t)
	named, okay := t.(*types.Named)

	if !okay {
		return false
	}

	o := named.Obj()

	if o.Pkg() == nil {
		return false
	}

	return o.Pkg().Path() == packagePath && o.Name() == name
}
