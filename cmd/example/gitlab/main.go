package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"github.com/funtimecoding/go-library/pkg/time"
	"github.com/spf13/pflag"
	"strings"
)

func main() {
	pflag.StringP(argument.Namespace, "n", "", "namespace")
	pflag.StringP(argument.Project, "p", "", "project")
	pflag.StringP(
		argument.Match,
		"m",
		"",
		"description match",
	)
	argument.ParseAndBind()
	g := gitlab.NewEnvironment()

	if true {
		runnerWay(g, argument.RequiredStringFlag(argument.Match, 1))
	}

	if false {
		p := g.ProjectByName(
			argument.RequiredStringFlag(argument.Namespace, 1),
			argument.RequiredStringFlag(argument.Project, 1),
		)

		if false {
			pipelineWay(g, p)
		}

		if false {
			projectWay(g, p)
		}
	}

	// TODO: Retry if failed, but only if not already retried and successful afterwards
}

func runnerWay(
	g *gitlab.Client,
	match string,
) {
	f := format.New().Color().Extended()
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
			fmt.Printf(
				"  Job: %s | %s | %s\n",
				job.CreatedAt.Format(time.DateMinute),
				job.Name,
				job.Status,
			)

			if job.Status == gitlab.Failed {
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

func projectWay(
	g *gitlab.Client,
	p *project.Project,
) {
	for _, job := range g.ProjectJobs(p.Identifier) {
		if job.Status != gitlab.Failed {
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

func pipelineWay(
	g *gitlab.Client,
	p *project.Project,
) {
	for _, pipeline := range g.Pipelines(p.Identifier) {
		fmt.Printf("Pipeline: %+v\n", pipeline.ID)

		for _, job := range g.PipelineJobs(p.Identifier, pipeline.ID) {
			fmt.Printf("  Job: %s | %s\n", job.Name, job.Status)
		}
	}
}
