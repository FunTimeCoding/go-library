package issue_enricher

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

type (
	OptionFunc  func(*Enricher)
	SliceResult func(*issue.Issue) []string
	FloatResult func(*issue.Issue) float64
)
