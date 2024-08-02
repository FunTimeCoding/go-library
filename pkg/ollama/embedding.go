package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) Embedding(
	model string,
	prompt string,
) []float64 {
	result, e := c.client.Embeddings(
		c.context,
		&api.EmbeddingRequest{Model: model, Prompt: prompt},
	)
	errors.PanicOnError(e)

	return result.Embedding
}
