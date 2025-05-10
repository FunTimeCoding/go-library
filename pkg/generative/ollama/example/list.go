package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
)

func List() {
	o := ollama.NewEnvironment()

	for _, m := range o.List() {
		fmt.Printf("Model: %+v\n", m)
	}
}
