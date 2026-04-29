package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) MustFileExists(
	p *project.Project,
	branch string,
	file string,
) bool {
	result, e := c.FileExists(p, branch, file)
	errors.PanicOnError(e)

	return result
}
