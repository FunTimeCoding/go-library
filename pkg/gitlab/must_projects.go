package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) MustProjects() []*project.Project {
	result, e := c.Projects()
	errors.PanicOnError(e)

	return result
}
