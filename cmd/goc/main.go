package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/ceph/goc"
)

func main() {
	argument.ParseAndBind()
	goc.Run(argument.Positional(0), false)
}
