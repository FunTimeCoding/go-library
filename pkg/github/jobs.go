package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/job"
)

func (c *Client) Jobs(
	owner string,
	name string,
	run int64,
) []*job.Job {
	result, _, e := c.client.Actions.ListWorkflowJobs(
		c.context,
		owner,
		name,
		run,
		nil,
	)
	errors.PanicOnError(e)

	return job.NewSlice(result.Jobs)
}
