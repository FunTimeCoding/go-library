package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/chroma"
)

func Collection() {
	c := chroma.NewEnvironment()

	for _, l := range c.Collections() {
		fmt.Printf("Collection: %s\n", l.Name())
	}
}
