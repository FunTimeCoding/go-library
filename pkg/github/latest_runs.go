package github

import (
	"github.com/funtimecoding/go-library/pkg/github/run"
	"github.com/google/go-github/v84/github"
)

func (c *Client) LatestRuns(owner, repo string) []*run.Run {
	page, r, e := c.client.Actions.ListRepositoryWorkflowRuns(
		c.context,
		owner,
		repo,
		&github.ListWorkflowRunsOptions{
			ListOptions: github.ListOptions{PerPage: 100},
		},
	)
	panicOnError(r, e)

	return run.NewSlice(page.WorkflowRuns)
}
