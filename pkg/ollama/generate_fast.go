package ollama

import (
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/ollama/generate"
	"github.com/funtimecoding/go-library/pkg/ollama/request"
)

func (c *Client) GenerateFast(prompt string) *generate.Generate {
	r := request.New(prompt)
	r.Model = constant.Llama321b

	return generate.New(c.Generate(r))
}
