package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
)

func main() {
	argument.ParseBind()
	n := argument.Positional(0)
	fmt.Printf("Positional argument 0: %s\n", n)
}
