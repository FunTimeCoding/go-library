package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/runner"

func (c *Client) enrichRunner(r *runner.Runner) *runner.Runner {
	if n := c.GraphRunner(
		r.Identifier,
	).Data.Runner.Managers.Nodes; len(n) > 0 {
		r.Address = n[0].IPAddress
	}

	r.Validate()

	return r
}
