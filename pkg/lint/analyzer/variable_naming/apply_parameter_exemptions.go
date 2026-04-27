package variable_naming

func applyParameterExemptions(variables []typedVariable) {
	typeCount := map[string]int{}

	for _, v := range variables {
		if v.kind != kindParameter {
			continue
		}

		typeCount[v.typ.String()]++
	}

	for i := range variables {
		v := &variables[i]

		if v.kind != kindParameter {
			continue
		}

		if isPrimitiveType(v.typ) {
			continue
		}

		if len(v.ident.Name) == 1 {
			continue
		}

		if typeCount[v.typ.String()] >= 2 {
			v.exempt = true

			continue
		}

		if isPrimitiveSlice(v.typ) {
			v.exempt = true
		}
	}
}
