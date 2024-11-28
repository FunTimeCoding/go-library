package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) RunnerJobs(runner int) []*gitlab.Job {
	result, _, e := c.client.Runners.ListRunnerJobs(
		runner,
		&gitlab.ListRunnerJobsOptions{
			ListOptions: gitlab.ListOptions{PerPage: PerPage100},
			OrderBy:     gitlab.Ptr(Identifier),
			Sort:        gitlab.Ptr(Descending),
		},
	)
	errors.PanicOnError(e)

	return result
}
