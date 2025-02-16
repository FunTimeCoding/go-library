package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func Issue() {
	c := sentry.NewEnvironment()

	for _, i := range c.AllIssues() {
		fmt.Printf("Issue: %s\n", i.Format(format.Plain))
	}

	if false {
		for _, o := range c.Organizations() {
			fmt.Printf("Organization: %s\n", o.Name)

			for _, p := range c.OrganizationProjects(o) {
				fmt.Printf("Project: %s\n", p.Name)

				for _, i := range c.Issues(
					o,
					p,
					sentry.PeriodFortnight,
				) {
					fmt.Printf("Issue: %s\n", i.Format(format.Plain))
				}
			}
		}
	}
}
