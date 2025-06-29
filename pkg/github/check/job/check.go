package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/check/job/option"
	"github.com/funtimecoding/go-library/pkg/github/constant"
)

func Check(o *option.Job) {
	c := github.NewEnvironment()

	if o.Notation && o.Verbose {
		// Verbose would break JSON stdout
		o.Verbose = false
	}

	jobs := c.FailedRuns(o.Verbose)

	if o.Notation {
		printNotation(jobs, o)

		return
	}

	owner := c.User().Name

	if o.Verbose {
		fmt.Printf("Owner: %s\n", owner)
	}

	f := constant.Format

	for _, r := range jobs {
		fmt.Printf("Run: %s\n", r.Format(f))
		fmt.Printf("%s fail\n", r.Repository().FullName)
	}
}
