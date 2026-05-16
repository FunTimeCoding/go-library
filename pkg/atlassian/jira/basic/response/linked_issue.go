package response

type LinkedIssue struct {
	Identifier string            `json:"id"`
	Key        string            `json:"key"`
	Self       string            `json:"self"`
	Fields     LinkedIssueFields `json:"fields"`
}
