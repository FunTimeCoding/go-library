package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func TrackedIssue() {
	for _, i := range sentry.NewEnvironment().TrackedIssues() {
		fmt.Printf("Issue: %s\n", i.Format(format.Plain))
	}
}
