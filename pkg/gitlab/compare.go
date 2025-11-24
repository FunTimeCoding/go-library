package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Compare(
	project int64,
	from, to string,
) *gitlab.Compare {
	result, r, e := c.client.Repositories.Compare(
		project,
		&gitlab.CompareOptions{
			From: ptr.To(from),
			To:   ptr.To(to),
		},
	)
	panicOnError(r, e)

	return result
}
