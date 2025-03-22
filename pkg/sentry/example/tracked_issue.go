package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func TrackedIssue() {
	for _, i := range sentry.NewEnvironment().TrackedIssues() {
		fmt.Printf("Issue: %s\n", i.Format(option.Plain))
	}
}
