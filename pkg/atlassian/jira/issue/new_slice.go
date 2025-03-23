package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
)

func NewSlice(
	value []jira.Issue,
	m *option.Issue,
) []*Issue {
	var result []*Issue

	for _, element := range value {
		result = append(result, New(&element, m))
	}

	return result
}
