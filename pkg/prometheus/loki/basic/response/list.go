package response

type List struct {
	Status string   `json:"status"`
	Labels []string `json:"data"`
}
