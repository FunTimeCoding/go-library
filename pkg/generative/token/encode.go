package token

import "github.com/funtimecoding/go-library/pkg/errors"

func Encode(s string) ([]uint, []string) {
	identifiers, tokens, e := Encoding("gpt-4o").Encode(
		s,
		nil,
		nil,
	)
	errors.PanicOnError(e)

	return identifiers, tokens
}
