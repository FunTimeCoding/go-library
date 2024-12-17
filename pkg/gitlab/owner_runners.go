package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) ownerRunners(o *gitlab.ListRunnersOptions) ([]*gitlab.Runner, *gitlab.Response) {
	t, r, e := c.client.Runners.ListRunners(o)
	errors.PanicOnError(e)

	return t, r
}
