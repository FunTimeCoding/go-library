package variable_naming

func filterSingleLetter(variables []typedVariable) []typedVariable {
	var result []typedVariable

	for _, v := range variables {
		if len(v.ident.Name) == 1 {
			result = append(result, v)
		}
	}

	return result
}
