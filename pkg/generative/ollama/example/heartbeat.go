package example

import "github.com/funtimecoding/go-library/pkg/generative/ollama"

func Heartbeat() {
	o := ollama.NewEnvironment()
	o.Heartbeat()
}
