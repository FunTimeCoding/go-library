package main

import "github.com/funtimecoding/go-library/pkg/pushover"

func main() {
	p := pushover.NewEnvironment()
	p.Send("test")
}
