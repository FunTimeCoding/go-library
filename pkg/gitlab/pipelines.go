package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Pipelines(project int) []*gitlab.PipelineInfo {
	result, _, e := c.client.Pipelines.ListProjectPipelines(
		project,
		&gitlab.ListProjectPipelinesOptions{
			ListOptions: gitlab.ListOptions{PerPage: PerPage1000},
		},
	)
	errors.PanicOnError(e)

	return result
}
