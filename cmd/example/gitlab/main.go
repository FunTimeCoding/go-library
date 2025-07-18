package main

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/check/clean_job"
	"github.com/funtimecoding/go-library/pkg/gitlab/example"
)

func main() {
	if true {
		clean_job.Check()
	}

	if false {
		example.Project()
		example.Search()
		example.Runner()
		example.GraphQuery()
		example.CloneAll()
	}
}
