package main

import "github.com/funtimecoding/go-library/pkg/kubernetes/example"

func main() {
	example.Job()

	if false {
		example.Event()
		example.Pod()
		example.Namespace()
		example.Node()
	}
}
