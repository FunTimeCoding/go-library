package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/time"
	"github.com/spf13/pflag"
	"strings"
)

func main() {
	pflag.StringP(argument.Namespace, "n", "", "namespace")
	pflag.StringP(argument.Project, "p", "", "project")
	argument.ParseAndBind()
	namespace := argument.RequiredStringFlag(argument.Namespace, 1)
	project := argument.RequiredStringFlag(argument.Project, 1)
	g := gitlab.NewEnvironment()
	f := format.New().Color().Extended()

	if true {
		runners := g.Runners(true)
		fmt.Printf("Runners (%d):\n", len(runners))

		for _, runner := range runners {
			fmt.Printf("  %s\n", runner.Format(f))

			for _, job := range g.RunnerJobs(runner.Identifier) {
				fmt.Printf(
					"  Job: %s | %s | %s\n",
					job.CreatedAt.Format(time.DateMinute),
					job.Name,
					job.Status,
				)

				if job.Status == "failed" {
					trace := g.Trace(job.Project.ID, job.ID)

					if strings.Contains(
						trace,
						"ERROR: Job failed (system failure): prepare environment: waiting for pod running: timed out waiting for pod to start.",
					) {
						fmt.Printf("    Trace: %s\n", trace)
					}
				}
			}
		}
	}

	p := g.ProjectByName(namespace, project)

	if false {
		for _, pipeline := range g.Pipelines(p.Identifier) {
			fmt.Printf("Pipeline: %+v\n", pipeline.ID)

			for _, job := range g.PipelineJobs(p.Identifier, pipeline.ID) {
				fmt.Printf("  Job: %s | %s\n", job.Name, job.Status)
			}
		}
	}

	if false {
		for _, job := range g.ProjectJobs(p.Identifier) {
			if job.Status != "failed" {
				continue
			}

			fmt.Printf(
				"Job: %d | %s | %s | %s | %s\n",
				job.ID,
				job.Name,
				job.CreatedAt.Format(time.DateMinute),
				job.Stage,
				job.Status,
			)

			fmt.Printf("Trace: %s\n", g.Trace(p.Identifier, job.ID))
		}
	}

	// TODO: Retry if failed, but only if not already retried and successful afterwards
}
