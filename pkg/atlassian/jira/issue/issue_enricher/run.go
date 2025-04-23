package issue_enricher

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func (e *Enricher) Run(v []*issue.Issue) []*issue.Issue {
	for _, l := range v {
		e.enrich(l)
	}

	return v
}
