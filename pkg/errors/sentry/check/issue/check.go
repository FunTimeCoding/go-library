package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/age_colorer"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
)

func Check(o *option.Issue) {
	c := sentry.NewEnvironment()
	elements := c.IssuesSimple(o.Verbose)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	if o.Copyable {
		f.Tag(tag.Copyable)
	}

	colorer := age_colorer.Default(elements...)

	for _, e := range elements {
		colorer.Set(e)
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(item.GoSentry.Plural)
	}
}
