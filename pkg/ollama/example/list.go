package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
)

func List() {
	o := ollama.NewEnvironment()

	for _, m := range o.List() {
		fmt.Printf("Model: %+v\n", m)
	}
}
