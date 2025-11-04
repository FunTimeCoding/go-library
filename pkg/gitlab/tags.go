package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Tags(project int) []*gitlab.Tag {
	result, r, e := c.client.Tags.ListTags(
		project,
		&gitlab.ListTagsOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
	)
	panicOnError(r, e)

	return result
}
