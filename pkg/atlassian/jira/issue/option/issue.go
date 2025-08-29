package option

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"

type Issue struct {
	Locator       string
	User          string
	WatchedIssues []string
	FieldMap      *field_map.Map
	Verbose       bool
	ClosedStatus  []string
}
