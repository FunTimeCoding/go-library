package ollama

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/ollama/generate_request"
	"github.com/funtimecoding/go-library/pkg/ollama/generate_response"
)

func (c *Client) GenerateNotation(prompt string) *generate_response.Response {
	r := generate_request.New(prompt)
	r.Format = notation.Marshall("json")

	return generate_response.New(c.Generate(r))
}
