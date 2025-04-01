package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
)

func Embed() {
	o := ollama.NewEnvironment()

	for _, element := range o.Embedding(
		constant.Llama31,
		"What are embeddings?",
	) {
		fmt.Printf("Embedding: %+v\n", element)
	}
}
