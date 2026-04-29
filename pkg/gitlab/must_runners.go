package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/runner"
)

func (c *Client) MustRunners(all bool) []*runner.Runner {
	result, e := c.Runners(all)
	errors.PanicOnError(e)

	return result
}
