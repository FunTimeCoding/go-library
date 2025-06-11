package profile

import "github.com/funtimecoding/go-library/pkg/notation"

func Parse(s string) []*Profile {
	var a []any
	notation.DecodeStrict(s, &a, false)
	var result []*Profile

	for _, a2 := range a[1].([]any) {
		a3 := a2.([]any)
		result = append(
			result,
			&Profile{Account: a3[2].(string), Email: a3[3].(string)},
		)
	}

	return result
}
