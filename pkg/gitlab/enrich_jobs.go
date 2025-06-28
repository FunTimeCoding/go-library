package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func enrichJobs(
	v []*job.Job,
	p *project.Project,
) []*job.Job {
	for _, j := range v {
		j.Project = p
	}

	return v
}
