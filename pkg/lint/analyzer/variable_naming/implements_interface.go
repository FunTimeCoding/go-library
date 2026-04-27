package variable_naming

import "go/types"

func implementsInterface(
	t types.Type,
	packagePath string,
	name string,
) bool {
	method, _, _ := types.LookupFieldOrMethod(t, true, nil, name)

	if method == nil {
		return false
	}

	_ = packagePath
	signature, okay := method.Type().(*types.Signature)

	if !okay {
		return false
	}

	switch name {
	case "Read":
		return signature.Params().Len() == 1 &&
			signature.Results().Len() == 2 &&
			isByteSlice(signature.Params().At(0).Type()) &&
			isBasicKind(signature.Results().At(0).Type(), types.Int) &&
			signature.Results().At(1).Type().String() == "error"
	case "Write":
		return signature.Params().Len() == 1 &&
			signature.Results().Len() == 2 &&
			isByteSlice(signature.Params().At(0).Type()) &&
			isBasicKind(signature.Results().At(0).Type(), types.Int) &&
			signature.Results().At(1).Type().String() == "error"
	}

	return false
}
