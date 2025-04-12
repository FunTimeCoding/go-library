package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/query"
)

func Print(o *option.Issue) {
	c := jira.NewEnvironment()

	if o.Notation {
		printNotation(c)

		return
	}

	f := constant.Format

	for _, i := range query.Open(c) {
		fmt.Println(i.Format(f))
	}
}
