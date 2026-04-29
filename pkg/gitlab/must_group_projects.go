package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) MustGroupProjects(identifier ...int64) []*project.Project {
	result, e := c.GroupProjects(identifier...)
	errors.PanicOnError(e)

	return result
}
