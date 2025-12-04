package main

import (
	git "github.com/funtimecoding/go-library/pkg/git/example"
	github "github.com/funtimecoding/go-library/pkg/github/example"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/clean_job"
	gitlab "github.com/funtimecoding/go-library/pkg/gitlab/example"
)

func main() {
	gitlab.BranchRequest()

	if false {
		git.BuildInformation()

		clean_job.Check()
		gitlab.Issue()
		gitlab.MergeRequest()
		gitlab.Feature()
		gitlab.File()
		gitlab.Project()
		gitlab.Search()
		gitlab.Runner()
		gitlab.GraphQuery()
		gitlab.CloneAll()

		github.PullRequest()
		github.Search()
		github.CleanJob()
	}
}
