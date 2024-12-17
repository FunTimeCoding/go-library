package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) ProjectJobs(project int) []*gitlab.Job {
	result, _, e := c.client.Jobs.ListProjectJobs(project, nil)
	errors.PanicOnError(e)

	return result
}
