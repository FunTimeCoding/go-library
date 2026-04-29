package ollama

import "github.com/ollama/ollama/api"

func (c *Client) List() ([]api.ListModelResponse, error) {
	r, e := c.client.List(c.context)

	if e != nil {
		return nil, e
	}

	return r.Models, nil
}
