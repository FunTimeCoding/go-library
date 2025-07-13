package option

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"
	"log"
)

func New(
	locator string,
	user string,
	watchedIssues []string,
	doneStatus []string,
	m *field_map.Map,
) *Issue {
	if len(doneStatus) == 0 {
		log.Panicf("done status cannot be empty")
	}

	return &Issue{
		Locator:       locator,
		User:          user,
		WatchedIssues: watchedIssues,
		DoneStatus:    doneStatus,
		FieldMap:      m,
	}
}
