package argument

type SearchIssueEvents struct {
	Identifier string `json:"identifier"`
	Query      string `json:"query"`
	Limit      int    `json:"limit"`
	Cursor     string `json:"cursor"`
}
