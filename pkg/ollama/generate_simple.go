package ollama

import (
	"github.com/funtimecoding/go-library/pkg/ollama/generate_request"
	"github.com/funtimecoding/go-library/pkg/ollama/generate_response"
)

func (c *Client) GenerateSimple(prompt string) *generate_response.Response {
	return generate_response.New(
		c.Generate(generate_request.New().Prompt(prompt).Get()),
	)
}
