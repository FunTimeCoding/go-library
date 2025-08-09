package validator

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func New(_ *issue.Issue) *Validator {
	return &Validator{}
}
