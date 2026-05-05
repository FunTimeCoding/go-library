package argument

type SearchEvents struct {
	Query   string `json:"query"`
	Project string `json:"project"`
	Limit   int    `json:"limit"`
	Cursor  string `json:"cursor"`
}
