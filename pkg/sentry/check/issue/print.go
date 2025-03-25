package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/age_colorer"
	statusOption "github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/face"
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
	var colorable []face.AgeColorable

	for _, i := range issues {
		colorable = append(colorable, i)
	}

	colorer := age_colorer.Default(colorable...)

	for _, i := range issues {
		colorer.Set(i)
		fmt.Println(i.Format(f))
	}

	if len(issues) == 0 {
		fmt.Println("No relevant issues")
	}
}
