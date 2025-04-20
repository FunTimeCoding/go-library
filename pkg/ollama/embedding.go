package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) Embedding(
	model string,
	p string,
) []float64 {
	result, e := c.client.Embeddings(
		c.context,
		&api.EmbeddingRequest{Model: model, Prompt: p},
	)
	errors.PanicOnError(e)

	return result.Embedding
}
