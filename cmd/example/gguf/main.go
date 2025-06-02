package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/gpustack/gguf-parser-go"
)

func main() {
	argument.ParseBind()
	result, e := gguf_parser.ParseGGUFFile(
		argument.RequiredPositional(0, "PATH", 1),
		gguf_parser.SkipLargeMetadata(),
		gguf_parser.UseMMap(),
	)
	errors.PanicOnError(e)
	fmt.Printf("Parameters: %+v\n", result.Header)
}
