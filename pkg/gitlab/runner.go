package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/runner"

func (c *Client) Runner(identifier int64) *runner.Runner {
	result, r, e := c.client.Runners.GetRunnerDetails(identifier)
	panicOnError(r, e)

	return c.enrichRunner(runner.FromDetail(result))
}
