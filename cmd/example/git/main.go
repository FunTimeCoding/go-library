package main

import (
	git "github.com/funtimecoding/go-library/pkg/git/example"
	github "github.com/funtimecoding/go-library/pkg/github/example"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/clean_job"
	gitlab "github.com/funtimecoding/go-library/pkg/gitlab/example"
)

func main() {
	if false {
		git.BuildInformation()
	}

	if false {
		clean_job.Check()
		gitlab.Project()
		gitlab.Search()
		gitlab.Runner()
		gitlab.GraphQuery()
		gitlab.CloneAll()
	}

	if false {
		github.Search()
		github.CleanJob()
	}
}
