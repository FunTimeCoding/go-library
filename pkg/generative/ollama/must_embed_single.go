package ollama

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustEmbedSingle(
	model string,
	v string,
) []float32 {
	result, e := c.EmbedSingle(model, v)
	errors.PanicOnError(e)

	return result
}
