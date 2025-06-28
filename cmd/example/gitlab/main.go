package main

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/check/clean_job"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/example"
)

func main() {
	if true {
		job.Check()
	}

	if false {
		example.Runner()
		example.GraphQuery()
		clean_job.Check()
		example.CloneAll()
	}
}
