package issue_enricher

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func (e *Enricher) enrich(i *issue.Issue) {
	if e.concerns != nil {
		i.SetConcerns(e.concerns(i))
	}

	if e.score != nil {
		i.SetScore(e.score(i))
	}
}
