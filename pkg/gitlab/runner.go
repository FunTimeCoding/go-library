package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/runner"
)

func (c *Client) Runner(identifier int) *runner.Runner {
	result, _, e := c.client.Runners.GetRunnerDetails(identifier)
	errors.PanicOnError(e)

	return c.enrichRunner(runner.FromDetail(result))
}
