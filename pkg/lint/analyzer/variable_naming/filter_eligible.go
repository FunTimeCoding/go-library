package variable_naming

func filterEligible(variables []typedVariable) []typedVariable {
	var result []typedVariable

	for _, v := range variables {
		if isEligible(v) {
			result = append(result, v)
		}
	}

	return result
}
