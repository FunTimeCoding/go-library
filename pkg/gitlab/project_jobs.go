package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) ProjectJobs(p *project.Project) []*job.Job {
	result, r, e := c.client.Jobs.ListProjectJobs(p.Identifier, nil)
	panicOnError(r, e)

	return c.enrichProjectJobs(job.NewSlice(result), p)
}
