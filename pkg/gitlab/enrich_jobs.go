package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/job"

func (c *Client) enrichJobs(v []*job.Job) []*job.Job {
	for _, j := range v {
		c.enrichJob(j)
	}

	return v
}
