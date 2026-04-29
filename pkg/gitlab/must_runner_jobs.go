package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
)

func (c *Client) MustRunnerJobs(
	runner int64,
	stopAfter int,
) []*job.Job {
	result, e := c.RunnerJobs(runner, stopAfter)
	errors.PanicOnError(e)

	return result
}
