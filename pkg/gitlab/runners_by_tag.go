package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/runner"
	"slices"
)

func (c *Client) RunnersByTag(tag string) []*runner.Runner {
	var result []*runner.Runner

	for _, r := range c.Runners(true) {
		r = c.Runner(r.Identifier)

		if slices.Contains(r.Tags, tag) {
			result = append(result, r)
		}
	}

	return c.enrichRunners(result)
}
