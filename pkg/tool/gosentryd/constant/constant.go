package constant

const (
	Name = "gosentryd"

	Whoami            = "whoami"
	FindOrganizations = "find_organizations"
	FindProjects      = "find_projects"
	FindReleases      = "find_releases"
	SearchIssues      = "search_issues"
	GetIssue          = "get_issue"
	GetIssueTagValues = "get_issue_tag_values"
	SearchIssueEvents = "search_issue_events"
	SearchEvents      = "search_events"
	UpdateIssue       = "update_issue"
	GetIssueEvent     = "get_issue_event"
	DeleteIssue       = "delete_issue"

	Unreachable = "Sentry API unreachable"

	ServerInstructions = "Error tracking and triage. Search unresolved issues, inspect stack traces and events, delete resolved errors. Check after deploying new service versions."
)
