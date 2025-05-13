package token

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/hupe1980/go-tiktoken"
)

func Encoding(model string) *tiktoken.Encoding {
	result, e := tiktoken.NewEncodingForModel(model)
	errors.PanicOnError(e)

	return result
}
