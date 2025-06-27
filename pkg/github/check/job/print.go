package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/check/job/option"
	"github.com/funtimecoding/go-library/pkg/github/constant"
)

func Print(o *option.Job) {
	c := github.NewEnvironment()
	f := constant.Format
	owner := c.User().Name

	if o.Notation {
		printNotation(c)

		return
	}

	if o.Verbose {
		fmt.Printf("Owner: %s\n", owner)
	}

	for _, r := range c.FailedRuns(o.Verbose) {
		fmt.Printf("Run: %s\n", r.Format(f))
		fmt.Printf("%s fail\n", r.Repository().FullName)
	}
}
