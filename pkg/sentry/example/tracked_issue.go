package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func TrackedIssue() {
	c := sentry.NewEnvironment()
	f := option.Color.Copy()

	for _, i := range c.TrackedIssues() {
		fmt.Printf("Issue: %s\n", i.Format(f))
	}
}
