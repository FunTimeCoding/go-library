package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
)

func Show() {
	o := ollama.NewEnvironment()
	fmt.Printf("Show: %+v\n", o.Show(constant.Llama31))
}
