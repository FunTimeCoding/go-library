package main

import "github.com/funtimecoding/go-library/pkg/kubernetes/example"

func main() {
	if true {
		example.Event()
	}

	if false {
		example.Pod()
		example.Namespace()
		example.Node()
	}
}
