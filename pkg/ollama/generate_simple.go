package ollama

import "github.com/ollama/ollama/api"

func (c *Client) GenerateSimple(prompt string) string {
	return c.Generate(
		&api.GenerateRequest{Prompt: prompt},
	).Response
}
