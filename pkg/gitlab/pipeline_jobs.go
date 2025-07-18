package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) PipelineJobs(
	project int,
	pipeline int,
) []*job.Job {
	result, _, e := c.client.Jobs.ListPipelineJobs(
		project,
		pipeline,
		&gitlab.ListJobsOptions{IncludeRetried: ptr.To[bool](true)},
	)
	errors.PanicOnError(e)

	return c.enrichJobs(job.NewSlice(result))
}
