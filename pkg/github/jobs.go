package github

import "github.com/funtimecoding/go-library/pkg/github/job"

func (c *Client) Jobs(
	owner string,
	name string,
	run int64,
) []*job.Job {
	result, r, e := c.client.Actions.ListWorkflowJobs(
		c.context,
		owner,
		name,
		run,
		nil,
	)
	panicOnError(r, e)

	return job.NewSlice(result.Jobs)
}
