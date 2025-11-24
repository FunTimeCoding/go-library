package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) PipelineJobs(
	project int64,
	pipeline int64,
) []*job.Job {
	result, r, e := c.client.Jobs.ListPipelineJobs(
		project,
		pipeline,
		&gitlab.ListJobsOptions{IncludeRetried: ptr.To[bool](true)},
	)
	panicOnError(r, e)

	return c.enrichJobs(job.NewSlice(result))
}
