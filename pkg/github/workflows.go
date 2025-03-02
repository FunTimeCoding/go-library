package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/workflow"
)

func (c *Client) Workflows(
	owner string,
	name string,
) []*workflow.Workflow {
	result, _, e := c.client.Actions.ListWorkflows(
		c.context,
		owner,
		name,
		nil,
	)
	errors.PanicOnError(e)

	return workflow.NewSlice(result.Workflows)
}
