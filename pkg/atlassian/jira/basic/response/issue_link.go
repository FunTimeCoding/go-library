package response

type IssueLink struct {
	Id           string        `json:"id"`
	Self         string        `json:"self"`
	Type         IssueLinkType `json:"type"`
	OutwardIssue LinkedIssue   `json:"outwardIssue"`
}
