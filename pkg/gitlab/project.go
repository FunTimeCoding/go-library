package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) Project(identifier int) *project.Project {
	result, _, e := c.client.Projects.GetProject(identifier, nil)
	errors.PanicOnError(e)

	return project.New(result)
}
