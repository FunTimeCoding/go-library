package job

import (
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/check/job/option"
	"github.com/funtimecoding/go-library/pkg/github/run"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func collect(
	c *github.Client,
	o *option.Job,
) []*run.Run {
	if o.Notation && o.Verbose {
		// Verbose breaks JSON stdout
		o.Verbose = false
	}

	return monitor.OnlyConcerns(c.Runs(true, o.Verbose), o.All)
}
