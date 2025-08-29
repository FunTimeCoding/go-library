package option

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"
	"log"
)

func New(
	locator string,
	user string,
	watchedIssues []string,
	closedStatus []string,
	m *field_map.Map,
) *Issue {
	if len(closedStatus) == 0 {
		log.Panicf("done status cannot be empty")
	}

	return &Issue{
		Locator:       locator,
		User:          user,
		WatchedIssues: watchedIssues,
		ClosedStatus:  closedStatus,
		FieldMap:      m,
	}
}
