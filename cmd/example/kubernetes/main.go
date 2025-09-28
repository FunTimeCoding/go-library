package main

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/job"
	"github.com/funtimecoding/go-library/pkg/kubernetes/example"
)

func main() {
	job.Run()

	if false {
		example.Event()
		example.Pod()
		example.Namespace()
		example.Node()
	}
}
