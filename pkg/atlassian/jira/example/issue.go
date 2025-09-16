package example

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/davecgh/go-spew/spew"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Issue() {
	j := internal.Jira()
	projectKey := environment.Get(constant.ProjectKeyEnvironment)
	issueType := issue.TaskType
	summary := "Stub summary"
	description := "Stub description"
	f := constant.Format

	if true {
		p := j.MetaProject(projectKey)
		fmt.Printf("Project: %s\n", p.Name)
		t := j.MetaIssueType(p, issueType)
		fmt.Printf("Issue type: %s\n", t.Name)

		for k, v := range j.IssueTypeFields(t) {
			fmt.Printf("Field: %s = %s\n", k, v)
		}

		r := j.NewIssue(p, projectKey, issueType, summary, description)
		fmt.Println("Prepared:")
		spew.Dump(r)

		if true {
			i := j.CreateNative(r)
			fmt.Println("Created:")
			fmt.Println(i.Format(f))
		}
	}

	if true {
		r := issue.RawStub()
		r.Fields.Reporter = j.User()
		r.Fields.Project = jira.Project{Key: projectKey}
		r.Fields.Type = jira.IssueType{Name: issueType}
		r.Fields.Summary = summary
		r.Fields.Description = description
		fmt.Println("Prepared:")
		spew.Dump(r)

		if false {
			i := j.CreateNative(r)
			fmt.Println("Created:")
			fmt.Println(i.Format(f))
		}
	}
}
