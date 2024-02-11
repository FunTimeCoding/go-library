package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) Tags(project int) []*gitlab.Tag {
	result, _, e := c.client.Tags.ListTags(
		project,
		&gitlab.ListTagsOptions{
			ListOptions: gitlab.ListOptions{PerPage: PerPage},
		},
	)
	errors.PanicOnError(e)

	return result
}
