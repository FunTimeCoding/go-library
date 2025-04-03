package main

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/example"
)

func main() {
	if true {
		job.Check()
	}

	if false {
		example.CloneAll()
	}
}
