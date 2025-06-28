package clean_job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func ProjectWay(
	g *gitlab.Client,
	p *project.Project,
	f *option.Format,
) {
	for _, j := range g.ProjectJobs(p) {
		if j.Status != constant.Failed {
			continue
		}

		fmt.Printf("Job: %s\n", j.Format(f))
		fmt.Printf("Trace: %s\n", g.Trace(p.Identifier, j.Identifier))
	}
}
