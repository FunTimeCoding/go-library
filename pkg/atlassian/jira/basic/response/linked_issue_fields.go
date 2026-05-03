package response

type LinkedIssueFields struct {
	Summary   string    `json:"summary"`
	Status    Status    `json:"status"`
	Priority  Priority  `json:"priority"`
	IssueType IssueType `json:"issuetype"`
}
