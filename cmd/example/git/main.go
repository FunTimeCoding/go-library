package main

import (
	git "github.com/funtimecoding/go-library/pkg/git/example"
	github "github.com/funtimecoding/go-library/pkg/github/example"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/clean_job"
	gitlab "github.com/funtimecoding/go-library/pkg/gitlab/example"
)

func main() {
	gitlab.Issue()
	github.PullRequest()
	gitlab.MergeRequest()

	if false {
		git.BuildInformation()

		clean_job.Check()
		gitlab.File()
		gitlab.Project()
		gitlab.Search()
		gitlab.Runner()
		gitlab.GraphQuery()
		gitlab.CloneAll()

		github.Search()
		github.CleanJob()
	}
}
