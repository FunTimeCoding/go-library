package issue

import "github.com/andygrunwald/go-jira"

func statusField(i *jira.Issue) string {
	if i.Fields != nil && i.Fields.Status != nil {
		return i.Fields.Status.Name
	}

	return ""
}
