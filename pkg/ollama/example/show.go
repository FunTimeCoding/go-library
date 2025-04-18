package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
)

func Show() {
	o := ollama.NewEnvironment()
	fmt.Printf("Show: %+v\n", o.Show(constant.Llama31))
}
