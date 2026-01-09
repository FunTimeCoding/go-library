package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) Embed(
	model string,
	v []string,
) [][]float32 {
	result, e := c.client.Embed(
		c.context,
		&api.EmbedRequest{Model: model, Input: v},
	)
	errors.PanicOnError(e)

	return result.Embeddings
}
