package main

import "github.com/funtimecoding/go-library/pkg/monitor/check/collect"

func main() {
	if true {
		collect.Loop()
	}

	if false {
		collect.Check()
	}
}
