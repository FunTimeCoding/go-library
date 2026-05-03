package response

type LinkedIssue struct {
	Id     string            `json:"id"`
	Key    string            `json:"key"`
	Self   string            `json:"self"`
	Fields LinkedIssueFields `json:"fields"`
}
