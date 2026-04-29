package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/run"
)

func (c *Client) MustProjectRuns(
	owner string,
	name string,
) []*run.Run {
	result, e := c.ProjectRuns(owner, name)
	errors.PanicOnError(e)

	return result
}
