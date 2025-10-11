package response

type Search struct {
	Issues        []*Issue `json:"issues"`
	NextPageToken string   `json:"nextPageToken"`
	IsLast        bool     `json:"isLast"`
}
