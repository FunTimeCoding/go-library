package validator

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func New(i *issue.Issue) *Validator {
	return &Validator{}
}
