package response

type Changelog struct {
	StartAt    int             `json:"startAt"`
	MaxResults int             `json:"maxResults"`
	Total      int             `json:"total"`
	Histories  []ChangeHistory `json:"histories"`
}
