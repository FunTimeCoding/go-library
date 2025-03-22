package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/sentry"
	"github.com/funtimecoding/go-library/pkg/sentry/check/issue/parameter"
)

func Print(p *parameter.Issue) {
	if p.Notation {
		printNotation()

		return
	}

	c := sentry.NewEnvironment()
	f := option.Color.Copy().Tag(tag.Link)
	issues := c.AllIssues()

	for _, i := range issues {
		fmt.Println(i.Format(f))
	}

	if len(issues) == 0 {
		fmt.Println("No issues")
	}
}
