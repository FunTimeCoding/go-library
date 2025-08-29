package issue

import "github.com/andygrunwald/go-jira"

func RawStub() *jira.Issue {
	return &jira.Issue{Fields: &jira.IssueFields{}}
}
