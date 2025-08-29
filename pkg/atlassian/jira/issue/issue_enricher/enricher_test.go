package issue_enricher

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"testing"
)

func TestEnricher(t *testing.T) {
	assert.NotNil(t, New())
	assert.NotNil(
		t,
		New(
			WithConcernFunction(
				func(i *issue.Issue) []string {
					return []string{}
				},
			),
			WithScoreFunction(
				func(*issue.Issue) float64 {
					return 0
				},
			),
			WithCommentNameFilter([]string{}),
		),
	)
}
