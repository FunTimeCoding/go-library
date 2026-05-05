package argument

type FindReleases struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
}
