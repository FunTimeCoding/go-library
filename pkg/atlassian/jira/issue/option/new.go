package option

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"

func New(
	locator string,
	user string,
	watchedIssues []string,
	m *field_map.Map,
) *Issue {
	return &Issue{
		Locator:       locator,
		User:          user,
		WatchedIssues: watchedIssues,
		FieldMap:      m,
	}
}
