package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/run"
	"github.com/google/go-github/v76/github"
)

func (c *Client) ProjectRuns(
	owner string,
	name string,
) []*run.Run {
	page, _, e := c.client.Actions.ListRepositoryWorkflowRuns(
		c.context,
		owner,
		name,
		&github.ListWorkflowRunsOptions{
			ListOptions: github.ListOptions{
				PerPage: 100, // Cannot go higher than 100
			},
		},
	)
	errors.PanicOnError(e)
	result := run.NewSlice(page.WorkflowRuns)

	for _, r := range result {
		r.Validate()
	}

	return result
}
