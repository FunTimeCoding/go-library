package clean_job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func PipelineWay(
	g *gitlab.Client,
	p *project.Project,
	f *option.Format,
) {
	for _, i := range g.Pipelines(p.Identifier) {
		fmt.Printf("Pipeline: %+v\n", i.ID)

		for _, j := range g.PipelineJobs(p.Identifier, i.ID) {
			fmt.Printf("  Job: %s\n", j.Format(f))
		}
	}
}
