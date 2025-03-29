package main

import "github.com/funtimecoding/go-library/pkg/ollama/example"

func main() {
	if true {
		example.Fast()
	}

	if false {
		example.Stream()
		example.Simple()
		example.Show()
		example.List()
		example.Embed()
		example.Running()
		example.Heartbeat()
	}
}
