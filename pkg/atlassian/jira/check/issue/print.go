package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/query"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Print(o *option.Issue) {
	c := jira.NewEnvironment()
	project := environment.Get(constant.ProjectEnvironment, 1)

	if o.Notation {
		printNotation(c, o)

		return
	}

	f := constant.Format

	for _, i := range query.Open(c, project) {
		fmt.Println(i.Format(f))
	}
}
