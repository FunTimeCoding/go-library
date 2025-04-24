package issue_enricher

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

type (
	SliceResult func(*issue.Issue) []string
	FloatResult func(*issue.Issue) float64
)
