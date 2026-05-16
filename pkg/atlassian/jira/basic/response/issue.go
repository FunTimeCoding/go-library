package response

type Issue struct {
	Expand     string      `json:"expand"`
	Identifier string      `json:"id"`
	Self       string      `json:"self"`
	Key        string      `json:"key"`
	Fields     IssueFields `json:"fields"`
	Changelog  *Changelog  `json:"changelog,omitempty"`
}
