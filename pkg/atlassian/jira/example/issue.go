package example

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/davecgh/go-spew/spew"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func Issue() {
	j := common.Jira()
	p := j.DefaultProjectKey()
	issueType := issue.TaskType
	summary := "Stub summary"
	description := "Stub description"
	f := constant.Format.Copy().Tag(tag.Link)
	var i *jira.Issue

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
