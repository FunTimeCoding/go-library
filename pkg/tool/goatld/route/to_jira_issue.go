package route

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"time"
)

func toJiraIssue(i *issue.Issue) server.JiraIssue {
	result := server.JiraIssue{
		Key:     i.Key,
		Summary: i.Summary,
		Status:  i.Status,
		Type:    i.Type,
	}

	if i.Description != "" {
		result.Description = &i.Description
	}

	if i.Priority != "" {
		result.Priority = &i.Priority
	}

	if i.Initials != "" {
		result.Assignee = &i.Initials
	}

	if len(i.Labels) > 0 {
		result.Labels = &i.Labels
	}

	if i.Create != nil && !i.Create.IsZero() {
		s := i.Create.Format(time.RFC3339)
		result.Created = &s
	}

	if i.Due != nil && !i.Due.IsZero() {
		s := i.Due.Format(time.RFC3339)
		result.Due = &s
	}

	if i.Link != "" {
		result.Link = &i.Link
	}

	return result
}
