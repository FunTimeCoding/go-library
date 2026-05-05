package convert

import "github.com/funtimecoding/go-library/pkg/tool/goatlassiand/generated/server"

type IssueWithCustomFields struct {
	*server.JiraIssue
	Links        []IssueLink       `json:"links,omitempty"`
	Comments     []IssueComment    `json:"comments,omitempty"`
	Checklist    []ChecklistItem   `json:"checklist,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}
