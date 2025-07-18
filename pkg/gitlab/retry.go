package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
)

func (c *Client) Retry(
	project int,
	jobIdentifier int,
) *job.Job {
	result, _, e := c.client.Jobs.RetryJob(project, jobIdentifier)
	errors.PanicOnError(e)

	return c.enrichJob(job.New(result))
}
