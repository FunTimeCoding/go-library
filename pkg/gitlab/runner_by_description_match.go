package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/runner"
	"strings"
)

func (c *Client) RunnerByDescriptionMatch(s string) *runner.Runner {
	for _, r := range c.Runners(true) {
		if strings.Contains(r.Description, s) {
			return r
		}
	}

	return nil
}
