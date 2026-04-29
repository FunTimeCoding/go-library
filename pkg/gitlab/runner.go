package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/runner"

func (c *Client) Runner(identifier int64) (*runner.Runner, error) {
	result, _, e := c.client.Runners.GetRunnerDetails(identifier)

	if e != nil {
		return nil, e
	}

	return c.enrichRunner(runner.NewDetail(result)), nil
}
