package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/github"
)

func Run(verbose bool) {
	c := github.NewEnvironment()
	f := option.ExtendedColor.Copy()
	owner := *c.User().Login

	if verbose {
		fmt.Printf("Owner: %s\n", owner)
	}

	for _, a := range c.ActionRepository() {
		repository := a.Name

		if verbose {
			fmt.Printf("Repository: %s/%s\n", owner, repository)
		}

		for i, r := range c.Runs(owner, repository) {
			if i > 0 {
				c.DeleteRun(owner, repository, *r.Raw.ID)

				continue
			}

			if verbose {
				fmt.Printf("Run %d: %s\n", i, r.Format(f))
			}

			for _, j := range c.Jobs(owner, repository, *r.Raw.ID) {
				if verbose {
					fmt.Printf("  Job: %s\n", j.Format(f))
				}

				if j.Fail() {
					fmt.Printf("%s fail\n", r.Repository().FullName)
				}
			}
		}
	}
}
