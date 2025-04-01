package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
)

func Running() {
	o := ollama.NewEnvironment()

	for _, element := range o.Running() {
		fmt.Printf("Running: %+v\n", element)
	}
}
