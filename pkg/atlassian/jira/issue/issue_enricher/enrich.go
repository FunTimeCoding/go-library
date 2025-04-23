package issue_enricher

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func (e *Enricher) enrich(i *issue.Issue) {
	if e.validator != nil {
		i.Concerns = e.validator(i)
	}
}
