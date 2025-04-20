package ollama

import (
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/ollama/generate_request"
	"github.com/funtimecoding/go-library/pkg/ollama/generate_response"
)

func (c *Client) GenerateFast(prompt string) *generate_response.Response {
	r := generate_request.New(prompt)
	r.Model = constant.Llama321b

	return generate_response.New(c.Generate(r))
}
