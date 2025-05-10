package ollama

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/generate_response"
)

func (c *Client) GenerateNotation(p string) *generate_response.Response {
	return c.Generate(generate_request.New().Prompt(p).Notation())
}
