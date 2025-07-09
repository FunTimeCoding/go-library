package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/age_colorer"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/sentry"
	"github.com/funtimecoding/go-library/pkg/sentry/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/sentry/constant"
)

func Print(o *option.Issue) {
	c := sentry.NewEnvironment()
	elements := c.IssuesSimple(o.Verbose)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format
	colorer := age_colorer.Default(elements...)

	for _, e := range elements {
		colorer.Set(e)
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
