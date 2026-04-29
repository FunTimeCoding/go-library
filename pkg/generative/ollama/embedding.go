package ollama

import "github.com/ollama/ollama/api"

func (c *Client) Embedding(
	model string,
	p string,
) ([]float64, error) {
	r, e := c.client.Embeddings(
		c.context,
		&api.EmbeddingRequest{Model: model, Prompt: p},
	)

	if e != nil {
		return nil, e
	}

	return r.Embedding, nil
}
