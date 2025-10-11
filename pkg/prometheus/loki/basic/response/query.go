package response

type Query struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}
