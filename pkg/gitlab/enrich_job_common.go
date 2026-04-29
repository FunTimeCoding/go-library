package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/job"

func (c *Client) enrichJobCommon(j *job.Job) {
	if j.Fail() {
		j.Trace = c.MustTrace(j.Project.Identifier, j.Identifier)
	}

	j.Validate()
}
