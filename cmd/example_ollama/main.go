package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
)

func main() {
	o := ollama.New()

	if false {
		fmt.Println(o.GenerateSimple("What is a car?"))
	}

	if true {
		fmt.Printf("Show: %+v\n", o.Show(constant.Llama31))
	}

	if true {
		for _, element := range o.List() {
			fmt.Printf("Model: %+v\n", element)
		}
	}

	if true {
		for _, element := range o.Running() {
			fmt.Printf("Running: %+v\n", element)
		}
	}

	if false {
		for _, element := range o.Embedding() {
			fmt.Printf("Embedding: %+v\n", element)
		}
	}

	o.Heartbeat()
}
