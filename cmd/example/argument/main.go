package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
)

func main() {
	a := argument.NewSimple("argument-example")
	a.ParseSimple()
	fmt.Printf(
		"Positional argument 0: %s\n",
		a.Argument(0),
	)
}
