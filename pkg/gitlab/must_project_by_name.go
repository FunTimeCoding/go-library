package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) MustProjectByName(
	namespace string,
	name string,
) *project.Project {
	result, e := c.ProjectByName(namespace, name)
	errors.PanicOnError(e)

	return result
}
