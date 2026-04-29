package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) MustProject(identifier int64) *project.Project {
	result, e := c.Project(identifier)
	errors.PanicOnError(e)

	return result
}
