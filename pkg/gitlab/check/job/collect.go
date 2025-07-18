package job

import (
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/gitlab/job"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func collect(o *option.Job) []*job.Job {
	return monitor.OnlyConcerns(
		gitlab.NewEnvironment().Jobs(o.Verbose),
		o.All,
	)
}
