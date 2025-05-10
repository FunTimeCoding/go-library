package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
)

func Simple() {
	o := ollama.NewEnvironment()
	fmt.Println(o.GenerateSimple("What is a car?").Text)
}
