package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Retry(
	project int,
	job int,
) *gitlab.Job {
	result, _, e := c.client.Jobs.RetryJob(project, job)
	errors.PanicOnError(e)

	return result
}
