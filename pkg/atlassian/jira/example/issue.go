package example

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/davecgh/go-spew/spew"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Issue() {
	j := internal.Jira()
	p := environment.Get(constant.DefaultProjectKeyEnvironment)
	issueType := issue.TaskType
	summary := "Stub summary"
	description := "Stub description"
	f := constant.Format.Copy().Tag(tag.Link)
	var i *jira.Issue

	if false {
		j.PrintIssueFields(p, issueType)
	}

	if true {
		i = j.NewIssue(p, issueType, summary, description)
	}

	if false {
		i = j.NewIssueUnverified(p, issueType, summary, description)
	}

	fmt.Println("Prepared:")
	spew.Dump(i)

	if false {
		fmt.Println("Created:")
		fmt.Println(j.CreateNative(i).Format(f))
	}
}
