package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
)

func main() {
	fmt.Println(ollama.New().GenerateSimple("What is a car?"))
}
