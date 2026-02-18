package clean_job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/runner"
)

func RunnerWay(
	g *gitlab.Client,
	r *runner.Runner,
	f *option.Format,
) {
	fmt.Printf("Runner: %s\n", r.Format(f))
	jobs := g.RunnerJobs(r.Identifier, 1000)
	fmt.Printf("Job count: %d\n", len(jobs))
	f2 := f.Copy().Extended()

	for _, j := range jobs {
		if j.Fail() {
			fmt.Println(j.Format(f2))
		} else {
			fmt.Println(j.Format(f))
		}

		if j.HasConcern(job.Timeout) {
			fmt.Printf("  Start timeout: %s\n", j.Stage)

			if console.AskConfirmation("Retry job?") {
				// TODO: Untested
				g.RetryJob(j)
			}
		}
	}
}
