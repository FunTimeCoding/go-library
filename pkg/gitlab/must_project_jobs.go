package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) MustProjectJobs(p *project.Project) []*job.Job {
	result, e := c.ProjectJobs(p)
	errors.PanicOnError(e)

	return result
}
