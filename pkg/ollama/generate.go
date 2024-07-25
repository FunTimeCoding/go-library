package ollama

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
	"k8s.io/utils/pointer"
)

func (c *Client) Generate(prompt string) string {
	var result string

	respFunc := func(r api.GenerateResponse) error {
		result = r.Response

		return nil
	}

	errors.PanicOnError(
		c.client.Generate(
			context.Background(),
			&api.GenerateRequest{
				Model:  "llama3.1",
				Stream: pointer.Bool(false),
				Prompt: prompt,
			},
			respFunc,
		),
	)

	return result
}
