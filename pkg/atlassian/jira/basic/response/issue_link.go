package response

type IssueLink struct {
	Identifier   string        `json:"id"`
	Self         string        `json:"self"`
	Type         IssueLinkType `json:"type"`
	OutwardIssue LinkedIssue   `json:"outwardIssue"`
}
