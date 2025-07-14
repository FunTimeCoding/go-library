package job

import (
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
)

func collect(o *option.Job) []*job.Job {
	g := gitlab.NewEnvironment()
	var result []*job.Job

	if o.All {
		result = g.Jobs(o.Verbose)
	} else {
		result = g.FailedJobs(o.Verbose)
	}

	return result
}
