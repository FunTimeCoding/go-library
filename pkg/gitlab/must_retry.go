package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
)

func (c *Client) MustRetry(
	project int64,
	jobIdentifier int64,
) *job.Job {
	result, e := c.Retry(project, jobIdentifier)
	errors.PanicOnError(e)

	return result
}
