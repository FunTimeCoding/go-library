package defer_close

import "go/types"

func implementsCloser(t types.Type) bool {
	closeMethod, _, _ := types.LookupFieldOrMethod(t, true, nil, "Close")

	if closeMethod == nil {
		return false
	}

	signature, okay := closeMethod.Type().(*types.Signature)

	if !okay {
		return false
	}

	if signature.Params().Len() != 0 {
		return false
	}

	if signature.Results().Len() != 1 {
		return false
	}

	return signature.Results().At(0).Type().String() == "error"
}
