package issue_enricher

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

type Validator func(*issue.Issue) []string
