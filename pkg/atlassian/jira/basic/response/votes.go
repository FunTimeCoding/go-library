package response

type Votes struct {
	Self     string `json:"self"`
	Votes    int    `json:"votes"`
	HasVoted bool   `json:"hasVoted"`
}
