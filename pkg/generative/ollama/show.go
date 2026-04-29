package ollama

import "github.com/ollama/ollama/api"

func (c *Client) Show(model string) (*api.ShowResponse, error) {
	return c.client.Show(c.context, &api.ShowRequest{Model: model})
}
