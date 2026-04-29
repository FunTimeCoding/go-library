package ollama

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustEmbedding(
	model string,
	p string,
) []float64 {
	result, e := c.Embedding(model, p)
	errors.PanicOnError(e)

	return result
}
