package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/xanzy/go-gitlab"
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
					PerPage: PerPage100,
					Page:    number,
				},
				OrderBy: gitlab.Ptr(Identifier),
				Sort:    gitlab.Ptr(Descending),
			},
		)
		errors.PanicOnError(e)
		result = append(result, page...)

		if len(page) < PerPage100 {
			break
		}

		if stopAfter > 0 && len(result) >= stopAfter {
			result = job.Deduplicate(result)
			fmt.Printf("Deduplicated: %d\n", len(result))

			if len(result) >= stopAfter {
				break
			}
		}

		number++
	}

	return result
}
