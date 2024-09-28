package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) allRunners(o *gitlab.ListRunnersOptions) ([]*gitlab.Runner, *gitlab.Response) {
	t, r, e := c.client.Runners.ListAllRunners(o)
	errors.PanicOnError(e)

	return t, r
}
