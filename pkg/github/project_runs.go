package github

import (
	"github.com/funtimecoding/go-library/pkg/github/run"
	"github.com/google/go-github/v83/github"
)

func (c *Client) ProjectRuns(
	owner string,
	name string,
) []*run.Run {
	page, r, e := c.client.Actions.ListRepositoryWorkflowRuns(
		c.context,
		owner,
		name,
		&github.ListWorkflowRunsOptions{
			ListOptions: github.ListOptions{
				PerPage: 100, // Cannot go higher than 100
			},
		},
	)
	panicOnError(r, e)
	result := run.NewSlice(page.WorkflowRuns)

	for _, s := range result {
		s.Validate()
	}

	return result
}
