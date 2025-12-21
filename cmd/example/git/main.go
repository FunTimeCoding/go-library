package main

import (
	git "github.com/funtimecoding/go-library/pkg/git/example"
	github "github.com/funtimecoding/go-library/pkg/github/example"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/clean_job"
	gitlab "github.com/funtimecoding/go-library/pkg/gitlab/example"
)

func main() {
	git.BuildInformation()

	if false {
		clean_job.Check()
		gitlab.BranchRequest()
		gitlab.CloneAll()
		gitlab.Feature()
		gitlab.File()
		gitlab.GraphQuery()
		gitlab.Issue()
		gitlab.MergeRequest()
		gitlab.Project()
		gitlab.Runner()
		gitlab.Search()

		github.BranchRequest()
		github.CleanJob()
		github.PullRequest()
		github.Search()
	}
}
