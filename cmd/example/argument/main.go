package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
)

func main() {
	argument.ParseBind()
	fmt.Printf(
		"Positional argument 0: %s\n",
		argument.Positional(0),
	)
}
