package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Tags(project int) []*gitlab.Tag {
	result, _, e := c.client.Tags.ListTags(
		project,
		&gitlab.ListTagsOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
	)
	errors.PanicOnError(e)

	return result
}
