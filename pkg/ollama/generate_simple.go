package ollama

import "github.com/funtimecoding/go-library/pkg/ollama/request"

func (c *Client) GenerateSimple(prompt string) string {
	return c.Generate(request.New(prompt)).Response
}
