package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) ProjectJobs(project int) []*gitlab.Job {
	result, _, e := c.client.Jobs.ListProjectJobs(project, nil)
	errors.PanicOnError(e)

	return result
}
