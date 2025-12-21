package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) CreateTag(
	project int64,
	name string,
	reference string,
	message string,
) *gitlab.Tag {
	result, r, e := c.client.Tags.CreateTag(
		project,
		&gitlab.CreateTagOptions{
			TagName: ptr.To(name),
			Ref:     ptr.To(reference),
			Message: ptr.To(message),
		},
	)
	panicOnError(r, e)

	return result
}
