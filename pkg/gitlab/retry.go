package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/job"

func (c *Client) Retry(
	project int64,
	jobIdentifier int64,
) *job.Job {
	result, r, e := c.client.Jobs.RetryJob(project, jobIdentifier)
	panicOnError(r, e)

	return c.enrichJob(job.New(result))
}
