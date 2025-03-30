package ollama

import (
	"github.com/funtimecoding/go-library/pkg/ollama/generate"
	"github.com/funtimecoding/go-library/pkg/ollama/request"
)

func (c *Client) GenerateSimple(prompt string) *generate.Generate {
	return generate.New(c.Generate(request.New(prompt)))
}
