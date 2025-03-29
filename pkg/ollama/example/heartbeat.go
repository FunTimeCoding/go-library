package example

import "github.com/funtimecoding/go-library/pkg/ollama"

func Heartbeat() {
	o := ollama.New()
	o.Heartbeat()
}
