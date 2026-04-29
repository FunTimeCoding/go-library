package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/run"
)

func (c *Client) MustLatestRuns(owner, repo string) []*run.Run {
	result, e := c.LatestRuns(owner, repo)
	errors.PanicOnError(e)

	return result
}
