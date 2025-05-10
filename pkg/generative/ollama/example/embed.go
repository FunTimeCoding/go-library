package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
)

func Embed() {
	o := ollama.NewEnvironment()

	for _, e := range o.Embedding(
		constant.Llama31,
		"What are embeddings?",
	) {
		fmt.Printf("Embedding: %+v\n", e)
	}
}
