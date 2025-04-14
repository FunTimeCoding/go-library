package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/age_colorer"
	statusOption "github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/sentry"
	"github.com/funtimecoding/go-library/pkg/sentry/check/issue/option"
)

func Print(o *option.Issue) {
	c := sentry.NewEnvironment()

	if o.Notation {
		printNotation(c)

		return
	}

	f := statusOption.Color.Copy().Tag(tag.Link)
	issues := c.IssuesSimple()
	colorer := age_colorer.Default(issues...)

	for _, i := range issues {
		colorer.Set(i)
		fmt.Println(i.Format(f))
	}

	if len(issues) == 0 {
		fmt.Println("No relevant issues")
	}
}
