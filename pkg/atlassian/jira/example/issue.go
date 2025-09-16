package example

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/davecgh/go-spew/spew"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"log"
)

func Issue() {
	j := internal.Jira()
	key := environment.Get(constant.ProjectKeyEnvironment)
	issueType := issue.TaskType
	summary := "Stub summary"
	description := "Stub description"
	f := constant.Format

	if true {
		p := j.CreateMeta(key).GetProjectWithKey(key)

		if p == nil {
			log.Panicf("project not found: %s", key)
		}

		fmt.Printf("Project: %s\n", p.Name)
		t := p.GetIssueTypeWithName(issueType)

		if t == nil {
			log.Panicf("issue type not found: %s", issueType)
		}

		fmt.Printf("Issue type: %s\n", t.Name)
		fields, fieldsFail := t.GetAllFields()
		errors.PanicOnError(fieldsFail)

		for k, v := range fields {
			fmt.Printf("Field: %s = %s\n", k, v)
		}

		r, e := jira.InitIssueWithMetaAndFields(
			p,
			p.GetIssueTypeWithName(issueType),
			map[string]string{
				"Project":     key,
				"Issue Type":  issueType,
				"Summary":     summary,
				"Description": description,
			},
		)
		errors.PanicOnError(e)
		fmt.Println("Prepared:")
		spew.Dump(r)

		if true {
			i := j.CreateNative(r)
			fmt.Printf("Created: %s\n", i.Format(f))
		}
	}

	if true {
		r := issue.RawStub()
		r.Fields.Reporter = j.User()
		r.Fields.Project = jira.Project{Key: key}
		r.Fields.Type = jira.IssueType{Name: issueType}
		r.Fields.Summary = summary
		r.Fields.Description = description
		fmt.Println("Prepared:")
		spew.Dump(r)

		if false {
			i := j.CreateNative(r)
			fmt.Printf("Created: %s\n", i.Format(f))
		}
	}
}
