package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/run"
)

func (c *Client) Runs(
	owner string,
	name string,
) []*run.Run {
	result, _, e := c.client.Actions.ListRepositoryWorkflowRuns(
		c.context,
		owner,
		name,
		nil,
	)
	errors.PanicOnError(e)

	return run.NewSlice(result.WorkflowRuns)
}
