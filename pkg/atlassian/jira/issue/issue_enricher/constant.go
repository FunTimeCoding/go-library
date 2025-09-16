package issue_enricher

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

type (
	Option func(*Enricher)
	Slice  func(*issue.Issue) []string
	Float  func(*issue.Issue) float64
)
