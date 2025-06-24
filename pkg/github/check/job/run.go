package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/job"
)

func Run() {
	c := github.NewEnvironment()
	f := option.ExtendedColor.Copy()
	owner := *c.User().Login

	// TODO: Filter which repositories most recent conclusion is failed
	for _, a := range c.ActionRepository() {
		repository := a.Name
		fmt.Printf("Repository: %s/%s\n", owner, repository)

		for _, r := range c.Runs(owner, repository) {
			fmt.Printf("Run: %s\n", r.Format(f))

			for _, j := range c.Jobs(owner, repository, *r.Raw.ID) {
				fmt.Printf("  Job: %s\n", j.Format(f))

				if false {
					fmt.Printf("    Conclusion: %+v\n", *j.Raw.Conclusion)
				}

				if *j.Raw.Conclusion == job.Failure {
					fmt.Printf("    Failed\n")
				}
			}

			if false {
				c.DeleteRun(owner, repository, *r.Raw.ID)
			}
		}
	}
}
