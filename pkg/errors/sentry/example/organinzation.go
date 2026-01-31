package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
)

func Organization() {
	c := sentry.NewEnvironment()

	for _, o := range c.Organizations() {
		fmt.Printf("Organization: %+v\n", o)

		for _, t := range c.OrganizationTeams(o) {
			fmt.Printf("Team: %+v\n", t)
		}
	}
}
