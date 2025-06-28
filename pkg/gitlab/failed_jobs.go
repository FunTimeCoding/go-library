package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/forge"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
)

func (c *Client) FailedJobs(verbose bool) []*job.Job {
	var result []*job.Job
	cleanup := forge.AutoCleanup()
	f := constant.Format

	for _, p := range c.PipelineProjects() {
		if verbose {
			fmt.Printf("Project: %s\n", p.Raw.NameWithNamespace)
		}

		for i, j := range c.ProjectJobs(p) {
			if i > 0 {
				if cleanup {
					c.DeletePipeline(p.Identifier, j.Raw.Pipeline.ID)
				}

				continue
			}

			if verbose {
				fmt.Printf("  Job: %s\n", j.Format(f))
			}

			if j.Fail() {
				result = append(result, j)
			}
		}
	}

	return result
}
