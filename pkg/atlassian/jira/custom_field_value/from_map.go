package custom_field_value

import "github.com/funtimecoding/go-library/pkg/notation"

func FromMap(m map[string]any) *Value {
	var result Value
	notation.DecodeStrict(
		notation.Encode(m, false),
		&result,
		true,
	)

	return &result
}
