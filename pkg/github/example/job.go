package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/github"
)

const (
	Owner      = "funtimecoding"
	Repository = "go-library"
)

func Job() {
	c := github.NewEnvironment()
	f := option.ExtendedColor.Copy()

	for _, w := range c.Workflows(Owner, Repository) {
		fmt.Printf("Workflow: %s\n", w.Format(f))
	}

	for _, r := range c.Runs(Owner, Repository) {
		fmt.Printf("Run: %s\n", r.Format(f))

		for _, j := range c.Jobs(Owner, Repository, *r.Raw.ID) {
			fmt.Printf("  Job: %s\n", j.Format(f))
		}

		if false {
			c.DeleteRun(Owner, Repository, *r.Raw.ID)
		}
	}
}
