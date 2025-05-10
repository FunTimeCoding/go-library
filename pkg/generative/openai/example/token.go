package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/hupe1980/go-tiktoken"
)

func Token() {
	encoding, e := tiktoken.NewEncodingForModel("gpt-4o")
	errors.PanicOnError(e)

	ids, tokens, e2 := encoding.Encode(
		"Explain a cat in 10 words.",
		nil,
		nil,
	)
	errors.PanicOnError(e2)

	fmt.Printf("IDs: %+v\n", ids)
	fmt.Printf("Tokens: %+v\n", tokens)
	fmt.Printf("Token count: %d\n", len(tokens))
}
