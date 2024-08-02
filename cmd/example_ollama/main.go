package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
)

func main() {
	// https://github.com/ollama/ollama/blob/main/docs/api.md
	o := ollama.New()

	if true {
		fmt.Println(o.GenerateSimple("What is a car?"))
	}

	if false {
		fmt.Printf("Show: %+v\n", o.Show(constant.Llama31))
	}

	if false {
		for _, element := range o.List() {
			fmt.Printf("Model: %+v\n", element)
		}
	}

	if false {
		for _, element := range o.Running() {
			fmt.Printf("Running: %+v\n", element)
		}
	}

	if false {
		for _, element := range o.Embedding(
			constant.Llama31,
			"What are embeddings?",
		) {
			fmt.Printf("Embedding: %+v\n", element)
		}
	}

	o.Heartbeat()
}
