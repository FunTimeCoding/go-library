package main

import (
	confluence "github.com/funtimecoding/go-library/pkg/atlassian/confluence/example"
	jira "github.com/funtimecoding/go-library/pkg/atlassian/jira/example"
	opsgenie "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/example"
)

func main() {
	confluence.Overview()

	if false {
		confluence.Export()
		confluence.Page()
		confluence.Search()
		confluence.Watch()
		confluence.Label()
		confluence.Search()
		confluence.Space()
		confluence.User()

		jira.Issue()
		jira.Search()
		jira.Customer()
		jira.Watch()
		jira.CustomValue()

		opsgenie.AddResponder()
		opsgenie.Alert()
	}
}
