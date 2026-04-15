package variable_naming

func filterEligible(variables []typedVariable) []typedVariable {
	var result []typedVariable

	for _, v := range variables {
		if len(v.ident.Name) == 1 {
			result = append(result, v)

			continue
		}

		if isErrorType(v.typ) {
			result = append(result, v)
		}
	}

	return result
}
