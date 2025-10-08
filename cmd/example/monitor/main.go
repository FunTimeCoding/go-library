package main

import "github.com/funtimecoding/go-library/pkg/monitor/check/collect"

func main() {
	collect.LoopIndividual()

	if false {
		collect.Loop()
		collect.Check(false, false)
	}
}
