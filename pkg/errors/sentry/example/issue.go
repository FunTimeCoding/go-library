package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
)

func Issue() {
	c := sentry.NewEnvironment()
	f := option.Color.Copy()

	for _, i := range c.IssuesSimple(true) {
		fmt.Printf("Issue: %s\n", i.Format(f))
	}

	if false {
		for _, o := range c.Organizations() {
			fmt.Printf("Organization: %s\n", o.Name)

			for _, p := range c.OrganizationProjects(o) {
				fmt.Printf("Project: %s\n", p.Name)

				for _, i := range c.Issues(
					o,
					p,
					constant.PeriodFortnight,
				) {
					fmt.Printf("Issue: %s\n", i.Format(f))
				}
			}
		}
	}
}
