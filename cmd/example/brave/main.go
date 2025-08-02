package main

import "github.com/funtimecoding/go-library/pkg/brave/example"

func main() {
	example.Send()

	if false {
		example.Open()
		example.Extract()
		example.Profile()
	}
}
