package variable_naming

import (
	"go/ast"
	"sort"
)

func assignLetters(variables []typedVariable) map[*ast.Ident]string {
	result := make(map[*ast.Ident]string)
	taken := make(map[string]bool)
	singleLetterVars := filterSingleLetter(variables)

	for _, v := range singleLetterVars {
		taken[v.ident.Name] = true
	}

	sort.Slice(
		singleLetterVars,
		func(i, j int) bool {
			return singleLetterVars[i].precedence < singleLetterVars[j].precedence
		},
	)
	taken = make(map[string]bool)

	for _, v := range singleLetterVars {
		letter := pickLetter(v.typ, taken)

		if letter == "" {
			continue
		}

		result[v.ident] = letter
		taken[letter] = true
	}

	return result
}
