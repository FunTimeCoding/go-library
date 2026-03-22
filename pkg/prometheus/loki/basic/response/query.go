package response

type Query struct {
	Status string      `json:"status"`
	Result QueryResult `json:"data"`
}
