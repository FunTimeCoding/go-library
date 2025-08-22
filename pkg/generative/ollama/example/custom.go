package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/generate_request"
)

func Custom() {
	model := constant.Llama31

	if true {
		//model = "llama2:7b"
		model = "codellama:7b"
	}

	o := ollama.NewEnvironment(ollama.WithSecure(true))
	fmt.Printf("Version: %s\n", o.Version())
	r := o.Generate(
		generate_request.New().Prompt(
			"One short sentence: What is a car?",
		).Model(model),
	)
	fmt.Println(r.Text)
	r.Print()
}
