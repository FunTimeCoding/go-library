package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	ProjectEnvironment    = "JIRA_PROJECT"
	ProjectKeyEnvironment = "JIRA_PROJECT_KEY"

	SearchLimit int = 100

	BasicSearchLimit int = 5000

	IssueCountType = "issueCount"

	AllowedValuesKey = "allowedValues"
)

// Field names
const (
	ProjectName     = "Project"
	IssueTypeName   = "Issue Type"
	SummaryName     = "Summary"
	DescriptionName = "Description"
)

const ParentEpic = "parentEpic" // Field

// Status
const (
	InProgress = "In Progress"
	Closed     = "Closed"

	ToDo = "To Do"
	Done = "Done"
)

var Format = option.ExtendedColor.Copy()

const (
	ServiceDeskClosed   = "Closed"
	ServiceDeskResolved = "Resolved"
)

var ServiceDeskDone = []string{
	ServiceDeskResolved,
	ServiceDeskClosed,
}
