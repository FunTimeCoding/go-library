package github

import "github.com/funtimecoding/go-library/pkg/github/workflow"

func (c *Client) Workflows(
	owner string,
	name string,
) []*workflow.Workflow {
	result, r, e := c.client.Actions.ListWorkflows(
		c.context,
		owner,
		name,
		nil,
	)
	panicOnError(r, e)

	return workflow.NewSlice(result.Workflows)
}
