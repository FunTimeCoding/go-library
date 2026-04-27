package variable_naming

func isEligible(v typedVariable) bool {
	if v.exempt {
		return false
	}

	if v.kind == kindReceiver {
		return true
	}

	if v.kind == kindParameter {
		if isPrimitiveType(v.typ) {
			return len(v.ident.Name) == 1
		}

		return true
	}

	if len(v.ident.Name) == 1 {
		return true
	}

	return isErrorType(v.typ)
}
