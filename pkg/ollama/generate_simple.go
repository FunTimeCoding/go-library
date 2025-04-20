package ollama

import (
	"github.com/funtimecoding/go-library/pkg/ollama/generate_request"
	"github.com/funtimecoding/go-library/pkg/ollama/generate_response"
)

func (c *Client) GenerateSimple(p string) *generate_response.Response {
	return c.Generate(generate_request.New().Prompt(p))
}
