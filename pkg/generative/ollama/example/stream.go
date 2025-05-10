package example

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/generate_response"
	"github.com/ollama/ollama/api"
)

func Stream() {
	c := make(chan string)
	r := api.GenerateResponse{}
	go ollama.NewEnvironment().GenerateStream(
		generate_request.New().Prompt(
			"One short sentence: What is a car?",
		).Get(),
		c,
		&r,
	)

	for output := range c {
		print(output)
	}

	println()
	generate_response.New(&r).Print()
}
