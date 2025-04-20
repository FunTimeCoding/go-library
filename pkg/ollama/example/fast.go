package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
)

func Fast() {
	o := ollama.NewEnvironment()
	r := o.GenerateFast("One short sentence: What is a car?")
	fmt.Println(r.Text)
	r.Format()
}
