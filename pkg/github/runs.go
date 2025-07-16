package github

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/forge"
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/github/run"
)

func (c *Client) Runs(
	loadJobs bool,
	verbose bool,
) []*run.Run {
	var result []*run.Run
	cleanup := forge.AutoCleanup()
	f := constant.Format
	owner := c.User().Name

	for _, a := range c.ActionRepository() {
		if verbose {
			fmt.Printf("Repository: %s/%s\n", owner, a.Name)
		}

		for i, r := range c.ProjectRuns(owner, a.Name) {
			if i > 0 {
				if cleanup {
					c.DeleteRun(owner, a.Name, r.Identifier)
				}

				continue
			}

			if verbose {
				fmt.Printf("Run %d: %s\n", i, r.Format(f))
			}

			if loadJobs {
				r.Jobs = c.Jobs(owner, a.Name, r.Identifier)
			}

			r.Validate()
			result = append(result, r)
		}
	}

	return result
}
