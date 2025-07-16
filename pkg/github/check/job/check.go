package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/check/job/option"
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Check(o *option.Job) {
	c := github.NewEnvironment()
	elements := collect(c, o)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	owner := c.User().Name

	if o.Verbose {
		fmt.Printf("Owner: %s\n", owner)
	}

	f := constant.Format.Copy().Tag(tag.Timestamp)

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
