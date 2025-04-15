package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
)

func NewSlice(
	v []jira.Issue,
	o *option.Issue,
) []*Issue {
	var result []*Issue

	for _, e := range v {
		result = append(result, New(&e, o))
	}

	return result
}
