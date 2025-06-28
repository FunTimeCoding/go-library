package clean_job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/time"
	"strings"
)

func RunnerWay(
	g *gitlab.Client,
	match string,
) {
	f := option.ExtendedColor.Copy()
	runners := g.Runners(true)
	fmt.Printf("Runners (%d):\n", len(runners))

	for _, runner := range runners {
		if !strings.Contains(runner.Description, match) {
			continue
		}

		fmt.Printf("  %s\n", runner.Format(f))
		jobs := g.RunnerJobs(runner.Identifier, 1000)
		fmt.Printf("  Jobs (%d):\n", len(jobs))

		for _, job := range jobs {
			var jobUser string

			if job.User != nil {
				jobUser = job.User.Name
			} else {
				jobUser = "no user"
			}

			fmt.Printf(
				"  Job: %s | %s | %s | %s\n",
				job.CreatedAt.Format(time.DateMinute),
				jobUser,
				job.Name,
				job.Status,
			)

			if job.Status == constant.Failed {
				trace := g.Trace(job.Project.ID, job.ID)

				if strings.Contains(
					trace,
					"ERROR: Job failed (system failure): prepare environment: waiting for pod running: timed out waiting for pod to start.",
				) {
					p := g.Project(job.Project.ID)
					fmt.Printf(
						"    Pod start timeout %s/%s in stage %s\n",
						p.Raw.Namespace.FullPath,
						p.Raw.Name,
						job.Stage,
					)

					if false {
						fmt.Printf("    Trace: %s\n", trace)
					}
				}
			}
		}
	}
}
