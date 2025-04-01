package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
)

func List() {
	o := ollama.NewEnvironment()

	for _, element := range o.List() {
		fmt.Printf("Model: %+v\n", element)
	}
}
