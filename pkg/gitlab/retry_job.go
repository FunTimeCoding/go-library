package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/job"

func (c *Client) RetryJob(j *job.Job) *job.Job {
	if j.Project == nil {
		panic("job has no project")
	}

	return c.Retry(j.Project.Identifier, j.Identifier)
}
