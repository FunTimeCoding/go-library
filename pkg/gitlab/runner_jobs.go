package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) RunnerJobs(
	runner int,
	stopAfter int,
) []*gitlab.Job {
	var result []*gitlab.Job
	var number int

	for {
		page, _, e := c.client.Runners.ListRunnerJobs(
			runner,
			&gitlab.ListRunnerJobsOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: constant.PerPage100,
					Page:    number,
				},
				OrderBy: gitlab.Ptr(constant.Identifier),
				Sort:    gitlab.Ptr(constant.Descending),
			},
		)
		errors.PanicOnError(e)
		result = append(result, page...)

		if len(page) < constant.PerPage100 {
			break
		}

		if stopAfter > 0 && len(result) >= stopAfter {
			result = job.Deduplicate(result)

			if len(result) >= stopAfter {
				break
			}
		}

		number++
	}

	return result
}
