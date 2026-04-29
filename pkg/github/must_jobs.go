package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/job"
)

func (c *Client) MustJobs(
	owner string,
	name string,
	run int64,
) []*job.Job {
	result, e := c.Jobs(owner, name, run)
	errors.PanicOnError(e)

	return result
}
