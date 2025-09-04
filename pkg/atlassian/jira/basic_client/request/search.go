package request

type Search struct {
	Expand          string   `json:"expand"`
	Fields          []string `json:"fields"`
	FieldsByKeys    bool     `json:"fieldsByKeys"`
	Jql             string   `json:"jql"`
	MaxResults      int      `json:"maxResults"`
	NextPageToken   string   `json:"nextPageToken"`
	Properties      []string `json:"properties"`
	ReconcileIssues []int    `json:"reconcileIssues"`
}
