package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
	"k8s.io/utils/ptr"
)

func (c *Client) PipelineJobs(
	project int,
	pipeline int,
) []*gitlab.Job {
	result, _, e := c.client.Jobs.ListPipelineJobs(
		project,
		pipeline,
		&gitlab.ListJobsOptions{IncludeRetried: ptr.To[bool](true)},
	)
	errors.PanicOnError(e)

	return result
}
