package main

import "github.com/funtimecoding/go-library/pkg/system/macos/example"

func main() {
	example.Bundle()

	if false {
		example.Sound()
		example.Notify()
	}
}
