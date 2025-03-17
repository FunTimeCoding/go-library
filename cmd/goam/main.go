package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/prometheus/amtool"
)

func main() {
	argument.ParseAndBind()
	amtool.Run(argument.Positional(0))
}
