package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) CreatePipeline(
	project int64,
	reference string,
	v []*gitlab.PipelineVariableOptions,
) *gitlab.Pipeline {
	result, r, e := c.client.Pipelines.CreatePipeline(
		project,
		&gitlab.CreatePipelineOptions{
			Ref:       ptr.To(reference),
			Variables: &v,
		},
	)
	panicOnError(r, e)

	return result
}
