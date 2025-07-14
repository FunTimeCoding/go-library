package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/runner"

func (c *Client) enrichRunner(r *runner.Runner) *runner.Runner {
	r.Validate()
	r.Address = c.GraphRunner(
		r.Identifier,
	).Data.Runner.Managers.Nodes[0].IPAddress

	return r
}
