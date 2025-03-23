package issue

import "github.com/andygrunwald/go-jira"

type Issue struct {
	Key         string
	Summary     string
	Description string

	Link string

	Raw *jira.Issue
}
