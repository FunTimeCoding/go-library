package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) ProjectJobs(p *project.Project) []*job.Job {
	result, _, e := c.client.Jobs.ListProjectJobs(p.Identifier, nil)
	errors.PanicOnError(e)

	return c.enrichProjectJobs(job.NewSlice(result), p)
}
