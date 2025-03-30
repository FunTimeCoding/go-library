package ollama

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/ollama/generate"
	"github.com/funtimecoding/go-library/pkg/ollama/request"
)

func (c *Client) GenerateNotation(prompt string) *generate.Generate {
	r := request.New(prompt)
	r.Format = notation.Marshall("json")

	return generate.New(c.Generate(r))
}
