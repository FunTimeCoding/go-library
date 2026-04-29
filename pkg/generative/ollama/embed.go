package ollama

import "github.com/ollama/ollama/api"

func (c *Client) Embed(
	model string,
	v []string,
) ([][]float32, error) {
	result, e := c.client.Embed(
		c.context,
		&api.EmbedRequest{Model: model, Input: v},
	)

	if e != nil {
		return nil, e
	}

	return result.Embeddings, nil
}
