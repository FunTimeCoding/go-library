package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
)

func (c *Client) MustRetryJob(j *job.Job) *job.Job {
	result, e := c.RetryJob(j)
	errors.PanicOnError(e)

	return result
}
