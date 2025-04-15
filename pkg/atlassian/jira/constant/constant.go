package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	ProjectEnvironment = "JIRA_PROJECT"

	SearchLimit int = 100

	IssueCountType = "issueCount"

	AllowedValuesKey = "allowedValues"
)

const ParentEpic = "parentEpic" // Field

const Closed = "Closed" // Status

var Format = option.ExtendedColor.Copy()

const (
	ServiceDeskClosed   = "Closed"
	ServiceDeskResolved = "Resolved"
)

var ServiceDeskDone = []string{
	ServiceDeskResolved,
	ServiceDeskClosed,
}
