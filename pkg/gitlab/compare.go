package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Compare(
	project int,
	from, to string,
) *gitlab.Compare {
	result, _, e := c.client.Repositories.Compare(
		project,
		&gitlab.CompareOptions{
			From: ptr.To(from),
			To:   ptr.To(to),
		},
	)
	errors.PanicOnError(e)

	return result
}
