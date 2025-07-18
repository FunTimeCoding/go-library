package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) enrichProjectJob(
	j *job.Job,
	p *project.Project,
) *job.Job {
	j.Project = p
	c.enrichJobCommon(j)

	return j
}
