package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/check/job/option"
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Check(o *option.Job) {
	c := github.NewEnvironment()

	if o.Notation && o.Verbose {
		// Verbose would break JSON stdout
		o.Verbose = false
	}

	elements := c.FailedRuns(o.Verbose)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	owner := c.User().Name

	if o.Verbose {
		fmt.Printf("Owner: %s\n", owner)
	}

	f := constant.Format

	for _, e := range elements {
		fmt.Printf("Run: %s\n", e.Format(f))
		fmt.Printf("%s fail\n", e.Repository().FullName)
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
