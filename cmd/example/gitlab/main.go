package main

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/example"
)

func main() {
	if true {
		example.Runner()
	}

	if false {
		example.GraphQuery()
		job.Check()
		example.CloneAll()
	}
}
