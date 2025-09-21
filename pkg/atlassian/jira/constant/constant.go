package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	DefaultProjectKeyEnvironment  = "JIRA_DEFAULT_PROJECT_KEY"
	DefaultProjectNameEnvironment = "JIRA_DEFAULT_PROJECT_NAME"
	DefaultIssueTypeEnvironment   = "JIRA_DEFAULT_ISSUE_TYPE"

	ClosedStatusEnvironment = "JIRA_CLOSED_STATUS"

	TestIssueEnvironment = "JIRA_TEST_ISSUE"
	TestFieldEnvironment = "JIRA_TEST_FIELD"

	SearchLimit int = 100

	BasicSearchLimit int = 5000

	IssueCountType = "issueCount"

	AllowedValuesKey = "allowedValues"
)

// Field names
const (
	AssigneeName     = "Assignee"
	AttachmentName   = "Attachment"
	DescriptionName  = "Description"
	IssueTypeName    = "Issue Type"
	LabelsName       = "Labels"
	LinkedIssuesName = "Linked Issues"
	ParentName       = "Parent"
	ProjectName      = "Project"
	ReporterName     = "Reporter"
	SummaryName      = "Summary"
	RankName         = "Rank"
	FlaggedName      = "Flagged"
	TeamName         = "Team"
	DevelopmentName  = "Development"
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
