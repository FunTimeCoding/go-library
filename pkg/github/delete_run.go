package github

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeleteRun(
	owner string,
	name string,
	run int64,
) {
	_, e := c.client.Actions.DeleteWorkflowRun(
		c.context,
		owner,
		name,
		run,
	)
	errors.PanicOnError(e)
}
