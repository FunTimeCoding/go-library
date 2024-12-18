package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"github.com/funtimecoding/go-library/pkg/time"
)

func ProjectWay(
	g *gitlab.Client,
	p *project.Project,
) {
	for _, job := range g.ProjectJobs(p.Identifier) {
		if job.Status != constant.Failed {
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
