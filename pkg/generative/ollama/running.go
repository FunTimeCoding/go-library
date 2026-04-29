package ollama

import "github.com/ollama/ollama/api"

func (c *Client) Running() ([]api.ProcessModelResponse, error) {
	r, e := c.client.ListRunning(c.context)

	if e != nil {
		return nil, e
	}

	return r.Models, nil
}
