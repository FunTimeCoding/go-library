package response

type Comment struct {
	Comments   []any  `json:"comments"`
	Self       string `json:"self"`
	MaxResults int    `json:"maxResults"`
	Total      int    `json:"total"`
	StartAt    int    `json:"startAt"`
}
