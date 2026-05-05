package argument

type GetIssueTagValues struct {
	Identifier string `json:"identifier"`
	Tag        string `json:"tag"`
	Limit      int    `json:"limit"`
}
