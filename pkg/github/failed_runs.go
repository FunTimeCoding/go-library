package github

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/forge"
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/github/run"
)

func (c *Client) FailedRuns(verbose bool) []*run.Run {
	var result []*run.Run
	cleanup := forge.AutoCleanup()
	f := constant.Format
	owner := c.User().Name

	for _, a := range c.ActionRepository() {
		name := a.Name

		if verbose {
			fmt.Printf("Repository: %s/%s\n", owner, name)
		}

		for i, r := range c.Runs(owner, name) {
			if i > 0 {
				if cleanup {
					c.DeleteRun(owner, name, r.Identifier)
				}

				continue
			}

			if verbose {
				fmt.Printf("Run %d: %s\n", i, r.Format(f))
			}

			for _, j := range c.Jobs(owner, name, r.Identifier) {
				if verbose {
					fmt.Printf("  Job: %s\n", j.Format(f))
				}

				if j.Fail() {
					result = append(result, r)
				}
			}
		}
	}

	return result
}
