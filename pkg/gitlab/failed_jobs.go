package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/job"

func (c *Client) FailedJobs(verbose bool) []*job.Job {
	var result []*job.Job

	for _, j := range c.Jobs(verbose) {
		if j.Fail() {
			result = append(result, j)
		}
	}

	return result
}
