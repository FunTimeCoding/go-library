package argument

type SearchIssues struct {
	Query   string `json:"query"`
	Project string `json:"project"`
	Limit   int    `json:"limit"`
	Cursor  string `json:"cursor"`
}
