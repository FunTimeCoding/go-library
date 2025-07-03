package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/query"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Check(o *option.Issue) {
	elements := query.Open(
		internal.Jira(),
		environment.Get(constant.ProjectEnvironment),
	)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
