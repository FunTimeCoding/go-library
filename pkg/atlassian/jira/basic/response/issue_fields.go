package response

import "github.com/trivago/tgo/tcontainer"

type IssueFields struct {
	StatusCategoryChangeDate      string           `json:"statuscategorychangedate"`
	FixVersions                   []any            `json:"fixVersions"`
	StatusCategory                StatusCategory   `json:"statusCategory"`
	Resolution                    any              `json:"resolution"`
	LastViewed                    any              `json:"lastViewed"`
	Priority                      Priority         `json:"priority"`
	Labels                        []any            `json:"labels"`
	TimeEstimate                  any              `json:"timeestimate"`
	AggregateTimeOriginalEstimate any              `json:"aggregatetimeoriginalestimate"`
	Versions                      []any            `json:"versions"`
	IssueLinks                    []IssueLink      `json:"issuelinks"`
	Assignee                      any              `json:"assignee"`
	Status                        Status           `json:"status"`
	Components                    []any            `json:"components"`
	AggregateTimeEstimate         any              `json:"aggregatetimeestimate"`
	Creator                       JiraUser         `json:"creator"`
	Subtasks                      []any            `json:"subtasks"`
	Reporter                      JiraUser         `json:"reporter"`
	AggregateProgress             Progress         `json:"aggregateprogress"`
	Progress                      Progress         `json:"progress"`
	Votes                         Votes            `json:"votes"`
	IssueType                     IssueType        `json:"issuetype"`
	TimeSpent                     any              `json:"timespent"`
	Project                       Project          `json:"project"`
	AggregateTimeSpent            any              `json:"aggregatetimespent"`
	ResolutionDate                any              `json:"resolutiondate"`
	WorkRatio                     int              `json:"workratio"`
	Watches                       Watches          `json:"watches"`
	IssueRestriction              IssueRestriction `json:"issuerestriction"`
	Created                       string           `json:"created"`
	Updated                       string           `json:"updated"`
	TimeOriginalEstimate          any              `json:"timeoriginalestimate"`
	Description                   any              `json:"description"`
	TimeTracking                  struct{}         `json:"timetracking"`
	Security                      any              `json:"security"`
	Attachment                    []any            `json:"attachment"`
	Summary                       string           `json:"summary"`
	Environment                   any              `json:"environment"`
	DueDate                       any              `json:"duedate"`
	Comment                       Comment          `json:"comment"`
	Unknowns                      tcontainer.MarshalMap
}
