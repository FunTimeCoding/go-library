package common

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/issue_enricher"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/validator"
)

func Jira() *jira.Client {
	return jira.NewEnvironment().Set(
		issue_enricher.New(
			issue_enricher.WithConcernFunction(
				func(i *issue.Issue) []string {
					return validator.New(i).Validate()
				},
			),
		),
	)
}

