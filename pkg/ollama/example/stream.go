package example

import (
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/ollama/generate"
	"github.com/funtimecoding/go-library/pkg/ollama/request"
	"github.com/ollama/ollama/api"
)

func Stream() {
	o := ollama.New()
	c := make(chan string)
	r := api.GenerateResponse{}
	go o.GenerateStream(
		request.New("One short sentence: What is a car?"),
		c,
		&r,
	)

	for output := range c {
		print(output)
	}

	println()
	generate.New(&r).Print()
}
