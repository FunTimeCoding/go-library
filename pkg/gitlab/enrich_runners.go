package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/runner"

func (c *Client) enrichRunners(v []*runner.Runner) []*runner.Runner {
	for _, e := range v {
		c.enrichRunner(e)
	}

	return v
}
