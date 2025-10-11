package response

type List struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}
