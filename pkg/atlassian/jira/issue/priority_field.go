package issue

import "github.com/andygrunwald/go-jira"

func priorityField(i *jira.Issue) string {
	if i.Fields != nil && i.Fields.Priority != nil {
		return i.Fields.Priority.Name
	}

	return DefaultPriority
}
