package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
)

func (c *Client) MustPipelineJobs(
	project int64,
	pipeline int64,
) []*job.Job {
	result, e := c.PipelineJobs(project, pipeline)
	errors.PanicOnError(e)

	return result
}
