package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Pipelines(project int64) []*gitlab.PipelineInfo {
	result, r, e := c.client.Pipelines.ListProjectPipelines(
		project,
		&gitlab.ListProjectPipelinesOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
	)
	panicOnError(r, e)

	return result
}
