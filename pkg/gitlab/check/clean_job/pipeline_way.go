package clean_job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func PipelineWay(
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
