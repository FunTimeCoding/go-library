package job

import (
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/check/job/option"
	"github.com/funtimecoding/go-library/pkg/github/run"
)

func collect(
	c *github.Client,
	o *option.Job,
) []*run.Run {
	if o.Notation && o.Verbose {
		// Verbose breaks JSON stdout
		o.Verbose = false
	}

	var result []*run.Run

	// TODO: Implement Validate() to check if a run is failed
	if o.All {
		result = c.Runs(true, o.Verbose)
	} else {
		result = c.FailedRuns(o.Verbose)
	}

	return result
}
