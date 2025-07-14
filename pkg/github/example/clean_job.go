package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/constant"
)

func CleanJob() {
	c := github.NewEnvironment()
	f := option.ExtendedColor.Copy()
	owner := constant.LibraryNamespace
	repository := constant.LibraryRepository

	for _, w := range c.Workflows(owner, repository) {
		fmt.Printf("Workflow: %s\n", w.Format(f))
	}

	for _, r := range c.ProjectRuns(owner, repository) {
		fmt.Printf("Run: %s\n", r.Format(f))

		for _, j := range c.Jobs(owner, repository, *r.Raw.ID) {
			fmt.Printf("  Job: %s\n", j.Format(f))
		}

		c.DeleteRun(owner, repository, *r.Raw.ID)
	}
}
