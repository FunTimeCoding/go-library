package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) RunnerJobs(
	runner int64,
	stopAfter int,
) []*job.Job {
	var result []*gitlab.Job
	var number int64

	for {
		page, r, e := c.client.Runners.ListRunnerJobs(
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
		panicOnError(r, e)
		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
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

	return c.enrichJobs(job.NewSlice(result))
}
