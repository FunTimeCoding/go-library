package variable_naming

func isEligible(v typedVariable) bool {
	if len(v.ident.Name) == 1 {
		return true
	}

	return isErrorType(v.typ)
}
