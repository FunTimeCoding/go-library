package main

import "github.com/funtimecoding/go-library/pkg/debian/example"

func main() {
	example.Download()

	if false {
		example.Netboot()
		example.Packer()
	}
}
