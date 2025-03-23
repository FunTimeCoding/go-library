package issue

import (
	"fmt"
	statusOption "github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/sentry"
	"github.com/funtimecoding/go-library/pkg/sentry/check/issue/option"
)

func Print(o *option.Issue) {
	if o.Notation {
		printNotation()

		return
	}

	c := sentry.NewEnvironment()
	f := statusOption.Color.Copy().Tag(tag.Link)
	issues := c.AllIssues()

	for _, i := range issues {
		fmt.Println(i.Format(f))
	}

	if len(issues) == 0 {
		fmt.Println("No relevant issues")
	}
}
